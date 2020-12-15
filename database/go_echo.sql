-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: go-mariadb:3306
-- Generation Time: Dec 15, 2020 at 08:51 PM
-- Server version: 10.4.3-MariaDB-1:10.4.3+maria~bionic
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_echo`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `action_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pin` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `verify` enum('waiting','yes','no') COLLATE utf8mb4_unicode_ci DEFAULT 'waiting',
  `verify_date` datetime DEFAULT NULL,
  `mobile` varchar(25) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` enum('owner','staff','other','admin','customer','brand-owner') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` enum('active','inactive','ban') COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_group_id` int(11) DEFAULT NULL,
  `gender` enum('male','female') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `name`, `last_name`, `email`, `remember_token`, `action_token`, `pin`, `verify`, `verify_date`, `mobile`, `type`, `status`, `user_group_id`, `gender`, `birthday`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'admin@admin.com', '$2y$10$NKE.pUShl4/JPqRpPgym1.fhXaumalps/lIrv2x0B6Iy9gM4GU0qG', 'Superadmin', 'administator', 'admin@admin.com', NULL, NULL, NULL, 'waiting', NULL, '1234567890', 'admin', 'active', 1, NULL, NULL, '2020-12-14 23:32:35', '2020-12-14 23:32:35', NULL),
(2, 'user@user.com', '$2y$10$bK94hCGGApJA.RLEWZc/quVVpwceL0hD0OblbXr2QOJo4aeldfIfu', 'User', '', 'user@user.com', NULL, NULL, NULL, 'waiting', NULL, '', NULL, 'active', 1, NULL, NULL, '2020-12-14 23:32:35', '2020-12-14 23:32:35', NULL),
(3, 'sirichai.jan@ascendcorp.com', '$2y$10$XEU4zOztvil8F5BKiCHMtueVVuFfMVhkg3MB3.36mxVB5AV2IJjSW', 'Dev Boy', 'Team P ches', 'sirichai.jan@ascendcorp.com', NULL, NULL, NULL, 'waiting', NULL, '1234567890', 'admin', 'active', 1, NULL, NULL, '2020-12-14 23:32:35', '2020-12-14 23:32:35', NULL);

-- --------------------------------------------------------