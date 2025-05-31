// sniper-start.js

// Terminal ekranını temizle
console.clear();

// Esprili mesaj listesi
const mesajlar = [
  "🔫 Sniper başlatıldı... Silinen mesajlar artık senden kaçamaz!",
  "🕵️‍♂️ Kod başladı ama mesajlar hâlâ korkudan silinmiyor.",
  "💡 MFA fix aktif! O sinir bozucu 2FA hatası artık yok.",
  "🎯 Hedef kitlendi: Silinen mesajlar CMD’ye düşecek.",
  "😎 Gözüm üstünüzde, mesaj sileni affetmem!",
  "📡 Discord sunucusu dinleniyor... Bir şey silindiğinde haberim olacak.",
  "⚠️ Uyarı: Sniper o kadar hızlı ki mesaj silen pişman!",
  "🚀 Sistem başlatıldı. Kod mükemmel, sen de öylesin (belki)."
];

// Rastgele mesaj seç
const rastgele = mesajlar[Math.floor(Math.random() * mesajlar.length)];

// Çıkış yazısı
console.log(`
════════════════════════════════════════════
👋 Selamlar! Sniper altyapısı aktif.

🛠️ İLK ÖNCELİKLE ANANI GOTTEN SIKTIK.
📦 BACININ AMCIGINI SİKTİM.
✅ RAT YEDİN EZİK ORUSPU EVLADI.
💓ANANI BURAK KAÇIRDI.
💗 VİOLENTİA SİKER.
🙀 LOGUN DÜŞTÜ ANANI SİKTİM.

🧠 Tavsiye: Sadece kendi hesabında test amaçlı kullan.

🎉 Mesaj: ${rastgele}
════════════════════════════════════════════
`);
