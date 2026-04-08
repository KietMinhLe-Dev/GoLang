-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Apr 08, 2026 at 02:46 AM
-- Server version: 8.2.0
-- PHP Version: 8.2.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `testcode`
--

-- --------------------------------------------------------

--
-- Table structure for table `appointments`
--

DROP TABLE IF EXISTS `appointments`;
CREATE TABLE IF NOT EXISTS `appointments` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `leadId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `createdById` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `status` enum('BOOKED','CANCELLED','NO_SHOW','SHOWED') DEFAULT NULL,
  `appointmentAt` timestamp NULL DEFAULT NULL,
  `branch` varchar(191) DEFAULT NULL,
  `hasCompanion` tinyint(1) DEFAULT NULL,
  `companionName` varchar(191) DEFAULT NULL,
  `companionPhone` varchar(50) DEFAULT NULL,
  `showDate` timestamp NULL DEFAULT NULL,
  `isClosed` tinyint(1) DEFAULT NULL,
  `serviceType` enum('MEMBER','PT','YOGA') DEFAULT NULL,
  `revenue` decimal(12,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tenantId` (`tenantId`),
  KEY `createdById` (`createdById`),
  KEY `idx_appointments_leadId` (`leadId`),
  KEY `idx_appointments_appointmentAt` (`appointmentAt`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `call_logs`
--

DROP TABLE IF EXISTS `call_logs`;
CREATE TABLE IF NOT EXISTS `call_logs` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `leadId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `userId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `status` enum('CONNECTED','NO_ANSWER','BUSY','UNREACHABLE','WRONG_NUMBER','CANCELLED') DEFAULT NULL,
  `cancelReason` varchar(191) DEFAULT NULL,
  `notes` varchar(191) DEFAULT NULL,
  `callid` varchar(191) DEFAULT NULL,
  `dnb` varchar(191) DEFAULT NULL,
  `ext` varchar(50) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `direction` varchar(50) DEFAULT NULL,
  `recordingUrl` varchar(191) DEFAULT NULL,
  `durationSecs` int DEFAULT '0',
  `calledAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `tenantId` (`tenantId`),
  KEY `idx_call_logs_leadId` (`leadId`),
  KEY `idx_call_logs_calledAt` (`calledAt`),
  KEY `idx_call_logs_userId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

DROP TABLE IF EXISTS `campaigns`;
CREATE TABLE IF NOT EXISTS `campaigns` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `name` varchar(191) DEFAULT NULL,
  `description` varchar(191) DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `tenantId` (`tenantId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `name`, `description`, `isActive`, `tenantId`, `createdAt`, `updatedAt`) VALUES
('792ee3a0-3263-11f1-aeb0-d843ae05a398', 'new', 'new', 1, '70af2187-3263-11f1-aeb0-d843ae05a398', '2026-04-07 09:23:44', '2026-04-07 09:23:44');

-- --------------------------------------------------------

--
-- Table structure for table `extentions`
--

DROP TABLE IF EXISTS `extentions`;
CREATE TABLE IF NOT EXISTS `extentions` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `ext` varchar(50) DEFAULT NULL,
  `password` varchar(191) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `isRecord` tinyint(1) DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `userId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tenantId` (`tenantId`),
  KEY `userId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `leads`
--

DROP TABLE IF EXISTS `leads`;
CREATE TABLE IF NOT EXISTS `leads` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `name` varchar(191) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `address` varchar(191) DEFAULT NULL,
  `source` varchar(191) DEFAULT NULL,
  `stage` enum('NEW','PROCESS','CONTACTED','BOOKED','SHOWED','CLOSED') DEFAULT NULL,
  `notes` varchar(191) DEFAULT NULL,
  `userId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `assignedAt` datetime DEFAULT NULL,
  `campaignId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `showAt` datetime DEFAULT NULL,
  `lastCallAt` datetime DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_leads_phone` (`phone`),
  KEY `idx_leads_userId` (`userId`),
  KEY `idx_leads_campaignId` (`campaignId`),
  KEY `idx_leads_stage` (`stage`),
  KEY `idx_leads_tenantId` (`tenantId`),
  KEY `idx_leads_lastCallAt` (`lastCallAt`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `leads`
--

INSERT INTO `leads` (`id`, `name`, `phone`, `address`, `source`, `stage`, `notes`, `userId`, `assignedAt`, `campaignId`, `tenantId`, `showAt`, `lastCallAt`, `createdAt`, `updatedAt`) VALUES
('7bf6fdb0-7ac0-49c6-acbe-fe2beeeb53ac', 'Nguyen Van A Updated', '0912345678', '45 Nguyen Hue, Q1, TP.HCM', 'TikTok', 'CONTACTED', 'Khach da nghe tu van', '2ec37d78-3263-11f1-aeb0-d843ae05a398', '2026-04-07 09:25:39', '792ee3a0-3263-11f1-aeb0-d843ae05a398', '70af2187-3263-11f1-aeb0-d843ae05a398', NULL, NULL, '2026-04-07 02:25:39', '2026-04-07 09:34:20'),
('d1754005-325f-11f1-aeb0-d843ae05a398', 'Lê Minh Kiệt', '0847733661', '337 Hồng Bàng', 'null\r\n', 'NEW', 'new', NULL, '2026-04-07 15:57:34', NULL, NULL, NULL, NULL, '2026-04-07 08:57:34', '2026-04-07 08:57:34');

-- --------------------------------------------------------

--
-- Table structure for table `notes`
--

DROP TABLE IF EXISTS `notes`;
CREATE TABLE IF NOT EXISTS `notes` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `content` varchar(191) DEFAULT NULL,
  `leadId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `createdById` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `leadId` (`leadId`),
  KEY `createdById` (`createdById`),
  KEY `tenantId` (`tenantId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tenants`
--

DROP TABLE IF EXISTS `tenants`;
CREATE TABLE IF NOT EXISTS `tenants` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `name` varchar(191) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `max_extensions` int DEFAULT NULL,
  `extensions` int DEFAULT NULL,
  `domain` varchar(191) DEFAULT NULL,
  `key_voicecloud` varchar(191) DEFAULT NULL,
  `pbx_url` varchar(191) DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `timezone` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `tenants`
--

INSERT INTO `tenants` (`id`, `name`, `status`, `max_extensions`, `extensions`, `domain`, `key_voicecloud`, `pbx_url`, `createdAt`, `updatedAt`, `timezone`) VALUES
('70af2187-3263-11f1-aeb0-d843ae05a398', 'new', '1', 1, 1, 'new', 'new', 'new', '2026-04-07 09:23:29', '2026-04-07 09:23:29', NULL),
('t1', 'Gym Center A', 'ACTIVE', 50, 10, 'gym-a.com', 'key123', 'pbx.gym-a.com', '2026-04-08 02:45:19', '2026-04-08 02:45:19', 'Asia/Ho_Chi_Minh');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `email` varchar(191) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `status` enum('ACTIVE','BUSY') DEFAULT NULL,
  `ext` varchar(50) DEFAULT NULL,
  `role` enum('0','1','2') DEFAULT NULL,
  `tenantId` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `tenantId` (`tenantId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `name`, `status`, `ext`, `role`, `tenantId`, `createdAt`, `updatedAt`) VALUES
('2ec37d78-3263-11f1-aeb0-d843ae05a398', 'new', 'new', 'ACTIVE', 'new', NULL, NULL, '2026-04-07 09:21:39', '2026-04-07 09:21:39'),
('a7f29a97-3263-11f1-aeb0-d843ae05a398', 'new@gmail.com', 'new', 'BUSY', 'new', '1', '70af2187-3263-11f1-aeb0-d843ae05a398', '2026-04-07 09:25:02', '2026-04-07 09:25:02');

--
-- Constraints for dumped tables
--

--
-- Constraints for table `appointments`
--
ALTER TABLE `appointments`
  ADD CONSTRAINT `appointments_ibfk_1` FOREIGN KEY (`leadId`) REFERENCES `leads` (`id`),
  ADD CONSTRAINT `appointments_ibfk_2` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`),
  ADD CONSTRAINT `appointments_ibfk_3` FOREIGN KEY (`createdById`) REFERENCES `users` (`id`);

--
-- Constraints for table `call_logs`
--
ALTER TABLE `call_logs`
  ADD CONSTRAINT `call_logs_ibfk_1` FOREIGN KEY (`leadId`) REFERENCES `leads` (`id`),
  ADD CONSTRAINT `call_logs_ibfk_2` FOREIGN KEY (`userId`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `call_logs_ibfk_3` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`);

--
-- Constraints for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD CONSTRAINT `campaigns_ibfk_1` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`);

--
-- Constraints for table `extentions`
--
ALTER TABLE `extentions`
  ADD CONSTRAINT `extentions_ibfk_1` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`),
  ADD CONSTRAINT `extentions_ibfk_2` FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

--
-- Constraints for table `leads`
--
ALTER TABLE `leads`
  ADD CONSTRAINT `leads_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `leads_ibfk_2` FOREIGN KEY (`campaignId`) REFERENCES `campaigns` (`id`),
  ADD CONSTRAINT `leads_ibfk_3` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`);

--
-- Constraints for table `notes`
--
ALTER TABLE `notes`
  ADD CONSTRAINT `notes_ibfk_1` FOREIGN KEY (`leadId`) REFERENCES `leads` (`id`),
  ADD CONSTRAINT `notes_ibfk_2` FOREIGN KEY (`createdById`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `notes_ibfk_3` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`);

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`tenantId`) REFERENCES `tenants` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
