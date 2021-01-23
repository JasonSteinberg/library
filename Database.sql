-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: library
-- ------------------------------------------------------
-- Server version	5.5.5-10.4.13-MariaDB-1:10.4.13+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `author`
--

DROP TABLE IF EXISTS `author`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `author` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `author`
--

INSERT INTO `author` VALUES (1,'Bill'),(2,'Eddie'),(4,'Jim'),(5,'Bob Mitchell'),(6,'Kinsey'),(7,'Kinsely'),(8,'Joe');

--
-- Table structure for table `book_author`
--

DROP TABLE IF EXISTS `book_author`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `book_author` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) DEFAULT NULL,
  `book_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book_author`
--

INSERT INTO `book_author` VALUES (5,1,1),(12,5,6),(20,7,7),(25,2,2),(26,1,2),(29,8,8);

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `description` varchar(4000) DEFAULT NULL,
  `isbn` varchar(20) DEFAULT NULL,
  `checkedout` int(11) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

INSERT INTO `books` VALUES (1,'Great Endings','Book of the year 2020','12312',1),(2,'Billings MT','Awesome','23112123',0),(6,'War and Peace','The longest Book ever','123123123',1),(7,'Web Dev','Best web dev book','123123123',1),(8,'Amazing Grace','Great Books','123123123',0);

--
-- Table structure for table `history`
--

DROP TABLE IF EXISTS `history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `history`
--

INSERT INTO `history` VALUES (1,1,1,'2021-01-23 01:52:56'),(2,1,0,'2021-01-23 01:53:02'),(3,1,1,'2021-01-23 01:53:10'),(4,1,0,'2021-01-23 01:53:13'),(5,1,1,'2021-01-23 01:54:05'),(6,1,0,'2021-01-23 01:54:09'),(7,1,1,'2021-01-23 01:54:42'),(8,1,0,'2021-01-23 01:55:04'),(9,1,1,'2021-01-23 01:55:11'),(10,1,0,'2021-01-23 02:13:33'),(11,1,1,'2021-01-23 02:13:45'),(12,1,1,'2021-01-23 18:54:23'),(13,1,0,'2021-01-23 18:55:18'),(14,1,1,'2021-01-23 18:55:55'),(15,1,0,'2021-01-23 18:56:54'),(16,1,1,'2021-01-23 18:57:49'),(17,2,0,'2021-01-23 18:57:52'),(18,2,1,'2021-01-23 18:58:56'),(19,3,1,'2021-01-23 18:59:33'),(20,7,1,'2021-01-23 19:00:41'),(21,6,1,'2021-01-23 19:02:11'),(22,2,0,'2021-01-23 19:03:59'),(23,2,1,'2021-01-23 19:04:21'),(24,3,0,'2021-01-23 19:04:24'),(25,3,1,'2021-01-23 19:04:25'),(26,2,0,'2021-01-23 19:05:48'),(27,3,0,'2021-01-23 19:05:50');