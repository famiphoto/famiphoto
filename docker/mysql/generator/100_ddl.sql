-- MySQL Script generated by MySQL Workbench
-- Sat Feb 15 23:53:16 2025
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema famiphoto_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema famiphoto_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `famiphoto_db` DEFAULT CHARACTER SET utf8 ;
USE `famiphoto_db` ;

-- -----------------------------------------------------
-- Table `famiphoto_db`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`users` (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT,
  `my_id` VARCHAR(128) NOT NULL,
  `status` INT NOT NULL,
  `is_admin` TINYINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `my_id_UNIQUE` (`my_id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`user_passwords`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`user_passwords` (
  `user_id` BIGINT NOT NULL,
  `password` VARCHAR(512) NOT NULL,
  `last_modified_at` DATETIME NOT NULL,
  `is_initialized` TINYINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_passwords_users_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `famiphoto_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photos` (
  `photo_id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(1024) NOT NULL,
  `imported_at` DATETIME NOT NULL,
  `description_ja` TEXT NOT NULL,
  `description_en` TEXT NOT NULL,
  `file_name_hash` VARCHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`photo_id`),
  INDEX `file_name_hash_idx` (`file_name_hash` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photo_files`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photo_files` (
  `photo_file_id` BIGINT NOT NULL AUTO_INCREMENT,
  `photo_id` BIGINT NOT NULL,
  `file_type` VARCHAR(45) NOT NULL,
  `file_path` TEXT NOT NULL,
  `file_hash` VARCHAR(128) NOT NULL,
  `file_path_hash` VARCHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`photo_file_id`),
  INDEX `fk_photo_files_photos_photo_id_idx` (`photo_id` ASC) VISIBLE,
  INDEX `file_hash_idx` (`file_hash` ASC) VISIBLE,
  INDEX `file_path_hash` (`file_path_hash` ASC) VISIBLE,
  CONSTRAINT `fk_photo_files_photos_photo_id`
    FOREIGN KEY (`photo_id`)
    REFERENCES `famiphoto_db`.`photos` (`photo_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photo_exif`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photo_exif` (
  `photo_exif_id` BIGINT NOT NULL AUTO_INCREMENT,
  `photo_id` BIGINT NOT NULL,
  `tag_id` INT NOT NULL,
  `tag_name` VARCHAR(512) NOT NULL,
  `tag_type` VARCHAR(128) NOT NULL,
  `value_string` TEXT NOT NULL,
  `sort_order` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`photo_exif_id`),
  INDEX `fk_photo_exif_photos_idx` (`photo_id` ASC) VISIBLE,
  CONSTRAINT `fk_photo_exif_photos`
    FOREIGN KEY (`photo_id`)
    REFERENCES `famiphoto_db`.`photos` (`photo_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
