-- --------------------------------------------------------
-- ホスト:                          127.0.0.1
-- サーバーのバージョン:                   5.7.21-log - MySQL Community Server (GPL)
-- サーバー OS:                      Win64
-- HeidiSQL バージョン:               9.5.0.5278
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- local_database のデータベース構造をダンプしています
CREATE DATABASE IF NOT EXISTS `local_database` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `local_database`;

--  テーブル local_database.twitters の構造をダンプしています
CREATE TABLE IF NOT EXISTS `twitters` (
  `ID` bigint(200) NOT NULL,
  `USER` text,
  `TWEETTEXT` longtext,
  `DATE` datetime DEFAULT NULL,
  `DATEINT` bigint(20) DEFAULT NULL,
  `Source` longtext,
  `mediaIDS` longtext,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- エクスポートするデータが選択されていません
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
