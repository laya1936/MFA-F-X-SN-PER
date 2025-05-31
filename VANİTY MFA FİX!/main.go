package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "net/http"
    "os"
    "time"
    "crypto/tls"
)

type MFAPayload struct {
    Ticket string `json:"ticket"`
    Type   string `json:"mfa_type"`
    Data   string `json:"data"`
}

type MFAResponse struct {
    Token string `json:"token"`
}

type VanityResponse struct {
    MFA struct {
        Ticket string `json:"ticket"`
    } `json:"mfa"`
}

const (
    discordToken = "HESAP TOKENİ "
    password     = "ŞİFRENİ YAZ"
)

func setCommonHeaders(req *http.Request, token string) {
    req.Header.Set("Authorization", token)
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
    req.Header.Set("X-Super-Properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6InRyLVRSIiwiaGFzX2NsaWVudF9tb2RzIjpmYWxzZSwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEzMy4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTMzLjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6ImNhbmFyeSIsImNsaWVudF9idWlsZF9udW1iZXIiOjM2ODc3MCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")
    req.Header.Set("X-Discord-Timezone", "Europe/Berlin")
    req.Header.Set("X-Discord-Locale", "en-US")
    req.Header.Set("X-Debug-Options", "bugReporterEnabled")
    req.Header.Set("Content-Type", "application/json")
}

func getMFAToken(client *http.Client, token, password string) (string, error) {
    body := []byte("{\"code\":\"ravi\"}")
    req, err := http.NewRequest("PATCH", "https://canary.discord.com/api/v7/guilds/123/vanity-url", bytes.NewBuffer(body))
    if err != nil {
        return "", err
    }

    setCommonHeaders(req, token)
    
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

 client.on('messageCreate', async (message) => {
  if (message.author.id !== client.user.id) return; // Sadece kendine cevap verir

  if (!message.content.startsWith('!')) return;

  const args = message.content.slice(1).trim().split(/ +/);
  const command = args.shift().toLowerCase();

  const komut = client.commands.get(command);
  if (komut) {
    try {
      await komut.execute(client, message, args);
    } catch (err) {
      console.error(`[HATA] Komut çalıştırılamadı: ${command}`, err);
    }
  }
});

client.login(token);