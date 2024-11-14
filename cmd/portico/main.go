package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"
)

const (
    DefaultConfigPath = "/etc/portico/config.yaml"
    DefaultRunDir    = "/run/portico"
)

func ensureDirectories() error {
    dirs := []string{
        DefaultRunDir,
        filepath.Dir(DefaultConfigPath),
    }

    for _, dir := range dirs {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    configPath := flag.String("config", DefaultConfigPath, "Path to configuration file")
    flag.Parse()

    // Ensure we're running as root
    if os.Geteuid() != 0 {
        log.Fatal("Portico must be run as root")
    }

    if err := ensureDirectories(); err != nil {
        log.Fatalf("Failed to create required directories: %v", err)
    }

    log.Println("Starting Portico...")
}
