USE [master]
GO
/****** Object:  Database [msbeer]    Script Date: 15/12/2021 23:48:40 ******/
CREATE DATABASE [msbeer]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'msbeer', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER\MSSQL\DATA\msbeer.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'msbeer_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER\MSSQL\DATA\msbeer_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
 WITH CATALOG_COLLATION = DATABASE_DEFAULT
GO
ALTER DATABASE [msbeer] SET COMPATIBILITY_LEVEL = 150
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [msbeer].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [msbeer] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [msbeer] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [msbeer] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [msbeer] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [msbeer] SET ARITHABORT OFF 
GO
ALTER DATABASE [msbeer] SET AUTO_CLOSE OFF 
GO
ALTER DATABASE [msbeer] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [msbeer] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [msbeer] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [msbeer] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [msbeer] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [msbeer] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [msbeer] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [msbeer] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [msbeer] SET  DISABLE_BROKER 
GO
ALTER DATABASE [msbeer] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [msbeer] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [msbeer] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [msbeer] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [msbeer] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [msbeer] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [msbeer] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [msbeer] SET RECOVERY FULL 
GO
ALTER DATABASE [msbeer] SET  MULTI_USER 
GO
ALTER DATABASE [msbeer] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [msbeer] SET DB_CHAINING OFF 
GO
ALTER DATABASE [msbeer] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [msbeer] SET TARGET_RECOVERY_TIME = 60 SECONDS 
GO
ALTER DATABASE [msbeer] SET DELAYED_DURABILITY = DISABLED 
GO
EXEC sys.sp_db_vardecimal_storage_format N'msbeer', N'ON'
GO
ALTER DATABASE [msbeer] SET QUERY_STORE = OFF
GO
USE [msbeer]
GO
/****** Object:  User [msbeer]    Script Date: 15/12/2021 23:48:40 ******/
CREATE USER [msbeer] WITHOUT LOGIN WITH DEFAULT_SCHEMA=[dbo]
GO
/****** Object:  User [ms_beer]    Script Date: 15/12/2021 23:48:40 ******/
CREATE USER [ms_beer] FOR LOGIN [ms_beer] WITH DEFAULT_SCHEMA=[dbo]
GO
/****** Object:  Table [dbo].[beer_item]    Script Date: 15/12/2021 23:48:40 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[beer_item](
	[ID] [int] NOT NULL,
	[Name] [varchar](10) NOT NULL,
	[Brewery] [varchar](10) NOT NULL,
	[Country] [varchar](10) NOT NULL,
	[Price] [numeric](18, 0) NOT NULL,
	[Currency] [varchar](10) NOT NULL
) ON [PRIMARY]
GO
USE [master]
GO
ALTER DATABASE [msbeer] SET  READ_WRITE 
GO
