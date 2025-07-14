-- 创建 goods 表
CREATE TABLE goods (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    number INT NOT NULL,
    img_url VARCHAR(255)
);

-- 插入 10 件商品数据
INSERT INTO goods (name, price, number, img_url) VALUES
('iPhone 14', 7999, 100, '/images/iphone14.jpg'),
('MacBook Pro', 12999, 50, '/images/macbook.jpg'),
('iPad Air', 4799, 80, '/images/ipadair.jpg'),
('Apple Watch', 2999, 120, '/images/applewatch.jpg'),
('AirPods Pro', 1799, 200, '/images/airpods.jpg'),
('Samsung Galaxy S23', 6999, 90, '/images/galaxys23.jpg'),
('Sony WH-1000XM5', 2799, 70, '/images/sonyaudio.jpg'),
('Dell XPS 13', 9999, 40, '/images/dellxps.jpg'),
('Microsoft Surface Pro 9', 8499, 60, '/images/surface.jpg'),
('Nintendo Switch OLED', 2099, 150, '/images/switch.jpg');    