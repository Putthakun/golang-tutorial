+ Map
  - Map = โครงสร้างข้อมูล key-value pair (เก็บข้อมูลเป็นคู่กุญแจ-ค่า)
  - Key ต้องมี type ที่ เปรียบเทียบได้ (comparable) เช่น string, int, bool (แต่ slice, map, function ใช้เป็น key ไม่ได้)

+ ถ้าสร้างแบบ nil จะต้องสร้าง make() ก่อนถึงจะกำหนด key-value ได้
