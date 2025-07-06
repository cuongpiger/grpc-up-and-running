package main

import (
    "context"
    "log"
    "net"
    "sync"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/reflection"
    grpcstt "google.golang.org/grpc/status"

    pb "github.com/cuongpiger/golang/healthcheck" // Adjust import path as needed
)

const (
    port = ":50051"
)

// HealthServer implements the Health service
type HealthServer struct {
    mu       sync.RWMutex
    services map[string]pb.HealthCheckResponse_ServingStatus
    watchers map[string][]chan pb.HealthCheckResponse_ServingStatus
    pb.UnimplementedHealthServer
}

// NewHealthServer creates a new health server instance
func NewHealthServer() *HealthServer {
    return &HealthServer{
        services: make(map[string]pb.HealthCheckResponse_ServingStatus),
        watchers: make(map[string][]chan pb.HealthCheckResponse_ServingStatus),
    }
}

// Check implements the Check method of the Health service
func (h *HealthServer) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    h.mu.RLock()
    defer h.mu.RUnlock()

    log.Printf("Health check requested for service: %s", req.Service)

    status, exists := h.services[req.Service]
    if !exists {
        log.Printf("Service %s not found", req.Service)
        return nil, grpcstt.Errorf(codes.NotFound, "service %s not found", req.Service)
    }

    log.Printf("Service %s status: %v", req.Service, status)
    return &pb.HealthCheckResponse{
        Status: status,
    }, nil
}

// Watch implements the Watch method of the Health service
func (h *HealthServer) Watch(req *pb.HealthCheckRequest, stream pb.Health_WatchServer) error {
    h.mu.Lock()
    
    // Create a channel for this watcher
    ch := make(chan pb.HealthCheckResponse_ServingStatus, 1)
    
    // Add to watchers list
    if h.watchers[req.Service] == nil {
        h.watchers[req.Service] = make([]chan pb.HealthCheckResponse_ServingStatus, 0)
    }
    h.watchers[req.Service] = append(h.watchers[req.Service], ch)
    
    // Send initial status
    currentStatus, exists := h.services[req.Service]
    if !exists {
        currentStatus = pb.HealthCheckResponse_UNKNOWN
    }
    
    h.mu.Unlock()
    
    log.Printf("Starting watch for service: %s", req.Service)
    
    // Send initial response
    if err := stream.Send(&pb.HealthCheckResponse{Status: currentStatus}); err != nil {
        h.removeWatcher(req.Service, ch)
        return err
    }

    // Listen for status changes
    for {
        select {
        case status := <-ch:
            log.Printf("Sending status update for service %s: %v", req.Service, status)
            if err := stream.Send(&pb.HealthCheckResponse{Status: status}); err != nil {
                h.removeWatcher(req.Service, ch)
                return err
            }
        case <-stream.Context().Done():
            log.Printf("Watch stream closed for service: %s", req.Service)
            h.removeWatcher(req.Service, ch)
            return stream.Context().Err()
        }
    }
}

// SetStatus sets the serving status for a service
func (h *HealthServer) SetStatus(service string, status pb.HealthCheckResponse_ServingStatus) {
    h.mu.Lock()
    defer h.mu.Unlock()

    log.Printf("Setting status for service %s to %v", service, status)
    h.services[service] = status

    // Notify all watchers
    for _, ch := range h.watchers[service] {
        select {
        case ch <- status:
        default:
            // Channel is full, skip this watcher
            log.Printf("Watcher channel full for service: %s", service)
        }
    }
}

// removeWatcher removes a watcher channel from the list
func (h *HealthServer) removeWatcher(service string, ch chan pb.HealthCheckResponse_ServingStatus) {
    h.mu.Lock()
    defer h.mu.Unlock()

    watchers := h.watchers[service]
    for i, watcher := range watchers {
        if watcher == ch {
            h.watchers[service] = append(watchers[:i], watchers[i+1:]...)
            close(ch)
            break
        }
    }
}

func main() {
    // Create health server
    healthServer := NewHealthServer()
    
    // Set initial health status
    healthServer.SetStatus("", pb.HealthCheckResponse_SERVING) // Overall server health
    healthServer.SetStatus("productinfo.ProductInfo", pb.HealthCheckResponse_SERVING)
    
    // Start gRPC server
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    
    // Register services
    pb.RegisterHealthServer(grpcServer, healthServer)
    
    // Enable reflection
    reflection.Register(grpcServer)

    log.Printf("🚀 gRPC Server with Health Check running on port %s", port)
    
    // Simulate status changes (for demonstration)
    go func() {
        time.Sleep(10 * time.Second)
        log.Println("Simulating service going down...")
        healthServer.SetStatus("productinfo.ProductInfo", pb.HealthCheckResponse_NOT_SERVING)
        
        time.Sleep(5 * time.Second)
        log.Println("Simulating service recovery...")
        healthServer.SetStatus("productinfo.ProductInfo", pb.HealthCheckResponse_SERVING)
    }()

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}