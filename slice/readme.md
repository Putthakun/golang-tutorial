
  + slice
    - slice เหมือน Array แต่ไม่จำเป็นต้องกำหนด index ตายตัวเหมือน Array
    - slice เป็น structure reference type
    - slice เก็บข้อมูล sequence 

    + capacity (cap) 
      - คือขนาดที่ slice จองไว้สำหรับจองสมาชิก มันจะเพิ่มขึ้นก้ต่อเมื่อมีการเพิ่ม  data    ลงไปเกินพื้นที่มันจองไว้ โดยจะเพิ่มขึ้นด้วยการคูณ 2 ไปเรื่อยๆ

    + lenght (len) 
      - คือจำนวนสมาชิกที่อยู่ใน slice

    + การเข้าถึง data
      - สามารถเข้าถึงเหมือน index ใน array

    + การ split data ใน slice
      - variableName[index_start:index_end]
      - โดยข้อมูลจะเริ่มแบ่งตั้งแต่ index_start จนถึง (index_end - 1)

    + **slice เป็น reference type ถ้าแก้ผ่าน slice ที่ copy มา ต้นฉบับก็จะโดนแก้ด้วย

    + nil คือการประกาศ slice โดยไม่ได้ allocate memory
      - var s []int
      - len cap = 0 
      - s == nil return true

    + empty allocate
      - empty := []int {}
      - len cap = 0

    + make()
      - make() เป็นฟังก์ชัน built-in ของ Go ที่ใช้ สร้างและ allocate memory สำหรับ slice, map, channel