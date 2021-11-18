/*
SQLyog Ultimate v12.5.1 (64 bit)
MySQL - 5.7.31 : Database - blog
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `blog`;

/*Table structure for table `article` */

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL COMMENT '用户名',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `cate_id` int(11) NOT NULL COMMENT '分类名称',
  `title` varchar(100) NOT NULL COMMENT '标题',
  `summary` varchar(255) NOT NULL COMMENT '摘要',
  `content` text NOT NULL COMMENT '内容',
  `view_count` int(10) DEFAULT '0' COMMENT '浏览量',
  `comment_count` int(10) DEFAULT '0' COMMENT '回复量',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1正常2禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

/*Data for the table `article` */

insert  into `article`(`id`,`user_name`,`email`,`cate_id`,`title`,`summary`,`content`,`view_count`,`comment_count`,`create_time`,`update_time`,`status`) values 
(1,'张三','111@qq.com',1,'aaaaa','bbbb','this is content',0,0,'2021-11-17 09:30:27',NULL,1),
(2,'张三','php@qq.com',2,'php','php.net','php is word best language',0,0,'2021-11-17 09:32:46',NULL,1),
(3,'phper','html@qq.com',2,'php','php.net','php is word best language',0,0,'2021-11-17 09:36:50',NULL,1),
(4,'htmler','php@qq.com',5,'html','https://www.w3school.com.cn/','<html>\n\n<head>\n  这里是文档的头部 ... ...\n  ...\n</head>\n\n<body>\n  这里是文档的主体 ... ...\n  ...\n</body>\n\n</html>',0,0,'2021-11-18 09:02:11',NULL,1),
(5,'javaer','java@qq.com',3,'java','https://www.java.com/zh-CN/','Oracle Java 许可重要更新\n从 2019 年 4 月 16 起的发行版更改了 Oracle Java 许可。\n新的适用于 Oracle Java SE 的 Oracle 技术网许可协议 与以前的 Oracle Java 许可有很大差异。 新许可允许某些免费使用（例如个人使用和开发使用），而根据以前的 Oracle Java 许可获得授权的其他使用可能会不再支持。 请在下载和使用此产品之前认真阅读条款。 可在此处查看常见问题解答。\n\n可以通过低成本的 Java SE 订阅 获得商业许可和技术支持。\n\nOracle 还在 jdk.java.net 的开源 GPL 许可下提供了最新的 OpenJDK 发行版。',0,0,'2021-11-18 09:04:09',NULL,1),
(6,'goer','go@qq.com',2,'go','https://studygolang.com/pkgdoc','2、promptui 0.9 发布\n\n命令行应用程序的交互式提示库。包括一些基于终端的优雅控件，例如密码输入，项目选择和确认提示。\n\n3、Pigo 1.4.5 发布\n\nPigo 是基于像素强度比较的物体检测纸张的纯 Go 脸部检测库，对人脸识别有兴趣的可以研究下。\n\n4、历史上的今天：Go 语言以开源方式向全球发布\n\n2009 年 11 月 10 日，Google 宣布发布 Go 语言。官方博文翻译文章：Go 12岁了！。',0,0,'2021-11-18 09:06:33',NULL,1);

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `cate_name` varchar(100) NOT NULL COMMENT '分类名称',
  `cate_sort` int(10) NOT NULL COMMENT '排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `category` */

insert  into `category`(`id`,`cate_name`,`cate_sort`,`create_time`,`update_time`) values 
(1,'php',9,'2021-11-15 11:09:49',NULL),
(2,'go',10,'2021-11-07 11:09:59',NULL),
(3,'java',8,'2021-11-01 11:10:19',NULL),
(4,'python',1,'2021-11-17 10:18:24',NULL),
(5,'html/css',10,'2021-11-17 10:20:07',NULL);

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `art_id` int(10) NOT NULL COMMENT '文章id',
  `content` text NOT NULL COMMENT '评论内容',
  `user_name` varchar(50) NOT NULL COMMENT '用户名',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态1正常2禁用',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `cteate_time` timestamp NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `comment` */

/*Table structure for table `leave` */

DROP TABLE IF EXISTS `leave`;

CREATE TABLE `leave` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '留言表',
  `user_name` varchar(50) NOT NULL COMMENT '用户名',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `content` text NOT NULL COMMENT '内容',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1正常2禁用',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `leave` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
