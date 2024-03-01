package main

import (
    "net/http"
    "os/exec"
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/disable-internet", func(c *gin.Context) {
        var json struct {
            Disable  bool `json:"disable"`
            Duration int  `json:"duration"` // Duration in minutes
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if json.Disable {
            cmd := exec.Command("netsh", "interface", "set", "interface", "Ethernet", "admin=DISABLED")
            err := cmd.Run()
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to disable internet"})
                return
            }

            if json.Duration > 0 {
                go func() {
                    time.Sleep(time.Duration(json.Duration) * time.Minute)
                    cmd := exec.Command("netsh", "interface", "set", "interface", "Ethernet", "admin=ENABLED")
                    err := cmd.Run()
                    if err != nil {
                        log.Printf("Failed to enable internet: %v", err)
                        return
                    }
                    log.Println("Internet enabled")
                }()
            }

            c.JSON(http.StatusOK, gin.H{"status": "Internet disabled", "enable_after_minutes": json.Duration})
        } else {
            c.JSON(http.StatusOK, gin.H{"status": "No action taken"})
        }
    })

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    router.POST("/shutdown", func(c *gin.Context) {
        var json struct {
            Shutdown bool `json:"shutdown"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if json.Shutdown {
            cmd := exec.Command("shutdown", "/s")
            err := cmd.Run()
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shutdown the system"})
                return
            }
            c.JSON(http.StatusOK, gin.H{"status": "System shutdown initiated"})
        } else {
            c.JSON(http.StatusOK, gin.H{"status": "No action taken"})
        }
    })

    router.Run(":8080")
}
