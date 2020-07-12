-- MySQL dump 10.14  Distrib 5.5.65-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: ramen_app
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `post_id` int DEFAULT NULL,
  `content` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,1,1,'おいしそう～'),(2,2,1,'通ってます～'),(3,1,10,'私も良く食べに行きます'),(4,1,1,'いいですね～'),(5,1,1,'お腹がすいてきたー\n早く食べたいですね');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `photo_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'テストログインユーザーあ','testuser@test.com','https://msp.c.yimg.jp/yjimage?q=GPugfsgXyLEtbvFl0k1rup2x8BjmM3COb2OSusjzsKIUWYIKz43JQKQ8xeF6m4UCX2Jnu5AsvfcK5Qy2peAm2d0k5sNaGI3P8_4zvrsChSfuhd0PoY4w1P_ZW3Ax4Zv3ML6gHa9t0dBU1edGEuZtjNI-&sig=13ct4akcu&x=297&y=170','$2a$10$wKA2J87iqo3DFRvwdTGRoeQH3dBdk8Ix8ImUylYQ/zeTLO6tto4Xa','2020-06-11 23:32:40','2020-06-11 23:32:40',NULL),(2,'ラーメン太郎','taro@test.com','http://res.cloudinary.com/dxo10noyu/raw/upload/v1592396338/pfqdr8n6usuhrj7hiond.jpg','$2a$10$jyQPRi56R8.AiwtfTkQY0.Lmsmd.ik/4vG9YbZMyBh3GjYi7E3USO','2020-06-11 23:36:33','2020-06-11 23:36:33',NULL),(3,'ジロリアン','jiro@test.com','https://msp.c.yimg.jp/yjimage?q=ZdXdH.sXyLG2xpRU9BcWwm5GEPCLmH9ZuNHX8clzVGezPWLLRv3ar5lhw8VNP8X5_gOSuoRMrET8IV3nLEI3CGAmhF65lietcAfXuS1PEaQYR2qfqd93pIYyhvNGagoJDHZU14xqt16brXSxVIJ1&sig=13al3m4mh&x=259&y=194','$2a$10$8cx45xk.XA6cTZ1lJqPc.eLWuRk1HjL3NHfk5eUPvW762YK6WTeQa','2020-06-11 23:37:49','2020-06-11 23:37:49',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-12 19:50:15
