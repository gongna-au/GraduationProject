drop database if exists `backend`;

create database `backend`;

use  `backend`;

-- 创建用户表
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(50) NOT NULL UNIQUE,
  `password` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `expenses` int(11) DEFAULT 0,
  `overage` int(11) DEFAULT 0,
  `email` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `postal_code` varchar(20) DEFAULT NULL,
  `about_me` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `users` (`phone`, `password`, `name`, `overage`, `email`, `address`, `city`, `country`, `postal_code`, `about_me`) VALUES
('12345678901', 'securePass1', 'Alice Johnson', 100, 'alice@example.com', '1234 Elm Street', 'Springfield', 'USA', '12345', 'Enthusiastic software engineer.'),
('12345678902', 'securePass2', 'Bob Smith', 200, 'bob@example.com', '2345 Maple Avenue', 'Lincoln', 'USA', '23456', 'Dedicated project manager.'),
('12345678903', 'securePass3', 'Carol White', 300, 'carol@example.com', '3456 Oak Lane', 'Madison', 'USA', '34567', 'Creative graphic designer.'),
('12345678904', 'securePass4', 'David Brown', 400, 'david@example.com', '4567 Pine Street', 'Jefferson', 'USA', '45678', 'Innovative marketing specialist.'),
('12345678905', 'securePass5', 'Eva Green', 500, 'eva@example.com', '5678 Birch Road', 'Clinton', 'USA', '56789', 'Passionate digital artist.'),
('12345678906', 'securePass6', 'Frank Wright', 600, 'frank@example.com', '6789 Cedar Blvd', 'Jackson', 'USA', '67890', 'Expert in cybersecurity.'),
('12345678907', 'securePass7', 'Grace Hall', 700, 'grace@example.com', '7890 Dogwood Ave', 'Hamilton', 'USA', '78901', 'Skilled software developer.'),
('12345678908', 'securePass8', 'Henry Adams', 800, 'henry@example.com', '8901 Elm St', 'Kingston', 'USA', '89012', 'Experienced business analyst.'),
('12345678909', 'securePass9', 'Ivy Wilson', 900, 'ivy@example.com', '9012 Maple Road', 'Lexington', 'USA', '90123', 'Talented UI/UX designer.'),
('12345678910', 'securePass10', 'Jake Martinez', 1000, 'jake@example.com', '1012 Oak Street', 'Monroe', 'USA', '01234', 'Dedicated team leader.'),
('12345678911', 'securePass11', 'Luna Clark', 1100, 'luna@example.com', '1113 Pine Avenue', 'Naperville', 'USA', '12345', 'Creative content writer.'),
('12345678912', 'securePass12', 'Miles Davis', 1200, 'miles@example.com', '1214 Birch Lane', 'Orlando', 'USA', '23456', 'Innovative product manager.'),
('12345678913', 'securePass13', 'Nora Lewis', 1300, 'nora@example.com', '1315 Cedar Street', 'Phoenix', 'USA', '34567', 'Expert data analyst.'),
('12345678914', 'securePass14', 'Oscar King', 1400, 'oscar@example.com', '1416 Dogwood Blvd', 'Queens', 'USA', '45678', 'Skilled network engineer.'),
('12345678915', 'securePass15', 'Piper Scott', 1500, 'piper@example.com', '1517 Elm Road', 'Raleigh', 'USA', '56789', 'Passionate product designer.'),
('15002597117', '123', 'Sayo Kravits', 1600, 'sayoKravits@example.com', '1618 Maple Ave', 'Stamford', 'USA', '67890', 'Enthusiastic operations manager.');

UPDATE `users`
SET `expenses` = 1300
WHERE `id` = 16;


Select * from `users`;

CREATE TABLE `bank_cards_bound` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `card_number` varchar(20) NOT NULL,
  `card_password` varchar(255) NOT NULL, -- 注意：实际应用中不应直接存储密码，而应存储密码的哈希值
  `bank_name` varchar(100) NOT NULL,
  `expiry_date`  varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `bank_cards_bound` (`user_id`, `card_number`, `card_password`, `bank_name`, `expiry_date`) VALUES 
(1, '4580967032145897', '784589', 'Global Bank', '2025-12-31'),
(1, '5523145801234567', '092388', 'City Finance Bank', '2026-11-30'),
(2, '4532157894561230', '667511', 'National Trust', '2027-10-31'),
(2, '4024007198765432', '889211', 'Pacific Bank', '2028-09-30'),
(3, '4916971532487601', '453111', 'Unity Finance', '2029-08-31'),
(3, '4556737586899855', '786211', 'Heritage Bank', '2030-07-31'),
(1, '4485637482910543', '112411', 'Prosperity Bank', '2031-06-30'),
(2, '4716351123456789', '334511', 'Metro Credit Union', '2032-05-31'),
(3, '4539871234567891', '556711', 'Highland Bank', '2033-04-30'),
(14, '4539871234567852', '785211', 'Highland Bank', '2033-04-30'),
(14, '4539871234565248', '524811', 'Highland Bank', '2033-04-30'),
(16, '4539871234567852', '785211', 'Highland Bank', '2033-04-30'),
(16, '4539871234565248', '524811', 'Highland Bank', '2033-04-30'),
(1, '4024007131234567', '778811', 'Crest Bank', '2034-03-31');


-- 创建停车场信息表
CREATE TABLE `parking_lots` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `imageUrl` VARCHAR(255) NOT NULL,
  `spaces` int(11) NOT NULL,
  `charge` DECIMAL(10, 2) NOT NULL, 
  `latitude` decimal(10, 8) NOT NULL,
  `longitude` decimal(11, 8) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `parking_lots` (`name`, `address`, `spaces`, `charge`, `latitude`, `longitude`, `imageUrl`) VALUES 
('Pike Place Market Parking Garage', '1531 Western Ave, Seattle, WA 98101 US', 90,  8.00, 47.609722, -122.342222, 'https://cdn.pixabay.com/photo/2014/04/02/10/45/parking-304465_1280.png'),
('South Place Market Parking Garage', '3231 Western Ave, Seattle, WA 98198 US', 100,  8.00, 47.609722, -122.332222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('West Place Market Parking Garage', '1356 Western Ave, Seattle, WA 12101 US', 80,  8.00, 47.609722, -122.322222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('East Place Market Parking Garage', '2356 Eastern Ave, Seattle, WA 98102 US', 85,  8.00, 47.609722, -122.312222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('North Place Market Parking Garage', '4356 Northern Ave, Seattle, WA 98103 US', 75,  8.00, 47.609722, -122.302222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('Central Place Market Parking Garage', '536 Central Ave, Seattle, WA 98104 US', 95,  8.00, 47.609722, -122.292222, 'https://cdn.pixabay.com/photo/2021/07/02/09/39/cars-6381364_1280.jpg'),
('Harbor Place Market Parking Garage', '636 Harbor Ave, Seattle, WA 98105 US', 85,  8.00, 47.609722, -122.282222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg'),
('River Place Market Parking Garage', '736 River Ave, Seattle, WA 98106 US', 65,  8.00, 47.609722, -122.272222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg'),
('Ocean Place Market Parking Garage', '836 Ocean Ave, Seattle, WA 98107 US', 55,  8.00, 47.609722, -122.262222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg'),
('Mountain Place Market Parking Garage', '936 Mountain Ave, Seattle, WA 98108 US', 45,  8.00, 47.609722, -122.252222, 'https://cdn.pixabay.com/photo/2016/03/22/00/45/multi-storey-car-park-1271917_1280.jpg');


-- 假设用户当前的经纬度为(latitude, longitude)
SET @user_lat = 47.609722;
SET @user_long = -122.342222;

SELECT *, 
       (6371 * acos(cos(radians(@user_lat)) 
                    * cos(radians(latitude)) 
                    * cos(radians(longitude) - radians(@user_long)) 
                    + sin(radians(@user_lat)) 
                    * sin(radians(latitude)))) AS distance 
FROM `parking_lots`
ORDER BY distance ASC;

CREATE TABLE `parking_history` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    parking_lot_id INT NOT NULL,
    vehicle_number VARCHAR(20) NOT NULL,
    payment_amount  DECIMAL(10, 2) NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL
);

INSERT INTO parking_history (user_id, parking_lot_id, start_time, end_time, vehicle_number,payment_amount) VALUES
(16, 8, '2024-02-24 07:45:00', '2024-02-24 08:30:00', 'EFI012',34.00),
(16, 4, '2024-02-25 07:45:00', '2024-02-26 08:30:00', 'YRI012',34.00),
(16, 8, '2024-02-26 07:45:00', '2024-02-27 08:30:00', 'EFI012',34.00),
(16, 4, '2024-02-27 07:45:00', '2024-02-28 08:30:00', 'YRI012',34.00),
(16, 8, '2024-02-28 07:45:00', '2024-02-29 08:30:00', 'EFI012',34.00),
(16, 4, '2024-03-01 07:45:00', '2024-03-02 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-03 07:45:00', '2024-03-04 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-05 07:45:00', '2024-03-06 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-07 07:45:00', '2024-03-08 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-09 07:45:00', '2024-03-10 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-11 07:45:00', '2024-03-12 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-13 07:45:00', '2024-03-14 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-15 07:45:00', '2024-03-16 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-17 07:45:00', '2024-03-18 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-19 07:45:00', '2024-03-20 07:45:00', 'YRI012',34.00),
(16, 4, '2024-03-21 07:45:00', '2024-03-22 07:45:00', 'YRI012',34.00),
(16, 4, '2024-02-23 07:45:00', '2024-02-24 07:45:00', 'YRI012',34.00);




CREATE TABLE `parking_reservations` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `parking_lot_id` INT NOT NULL,
  `vehicle_number` VARCHAR(20) NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  `status`VARCHAR(20) DEFAULT 'PENDING',
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`parking_lot_id`) REFERENCES `parking_lots`(`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



INSERT INTO `parking_reservations` (`user_id`, `parking_lot_id`, `vehicle_number`,`start_time`, `end_time`) VALUES
(1, 1, 'ASD430','2023-03-01 08:00:00', '2023-03-01 12:00:00'),
(2, 2, 'ASD135','2023-03-02 09:00:00', '2023-03-02 11:00:00'),
(3, 3, 'ASD235','2023-03-03 10:00:00', '2023-03-03 14:00:00'),
(1, 4, 'ASD335','2023-03-04 08:30:00', '2023-03-04 10:30:00'),
(2, 5, 'ASD435','2023-03-05 09:30:00', '2023-03-05 13:30:00'),
(3, 1, 'ASD435','2023-03-06 10:30:00', '2023-03-06 12:30:00'),
(1, 2, 'ASD635','2023-03-07 11:00:00', '2023-03-07 15:00:00'),
(2, 3, 'ASD735','2023-03-08 12:00:00', '2023-03-08 16:00:00'),
(3, 4, 'ASD835','2023-03-09 13:00:00', '2023-03-09 17:00:00'),
(14, 3, 'ASD935','2023-03-05 12:00:00', '2023-03-05 16:00:00'),
(14, 3, 'ASD835','2023-03-06 12:00:00', '2023-03-06 16:00:00'),
(14, 3, 'ASD735','2023-03-08 12:00:00', '2023-03-08 16:00:00'),
(14, 4, 'ADD435','2023-03-09 13:00:00', '2023-03-09 17:00:00'),
(16, 3, 'AQD435','2023-03-10 12:00:00', '2023-03-11 16:00:00'),
(16, 3, 'AAD435','2023-03-11 12:00:00', '2023-03-12 16:00:00'),
(16, 4, 'ARD435','2023-03-12 13:00:00', '2023-03-13 17:00:00'),
(1, 5, 'AED435','2023-03-12 14:00:00', '2023-03-13 18:00:00');


CREATE TABLE completed_parking_transactions (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  parking_lot_id INT NOT NULL,
  vehicle_number VARCHAR(20) NOT NULL,
  start_time DATETIME NOT NULL,
  end_time DATETIME NOT NULL,
  payment_amount DECIMAL(10, 2) NOT NULL,
  payment_method VARCHAR(50),  -- 如 'Credit Card', 'Cash', 'Online Payment' 等
  reservation_id INT,  -- 若基于预定，则存储预定的ID；否则为NULL
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (parking_lot_id) REFERENCES parking_lots(id),
  FOREIGN KEY (reservation_id) REFERENCES parking_reservations(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
