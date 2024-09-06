package main

// Phpmyadmin - Web interface for MySQL and MariaDB
// Homepage: https://www.phpmyadmin.net

import (
	"fmt"
	
	"os/exec"
)

func installPhpmyadmin() {
	// Método 1: Descargar y extraer .tar.gz
	phpmyadmin_tar_url := "https://files.phpmyadmin.net/phpMyAdmin/5.2.1/phpMyAdmin-5.2.1-all-languages.tar.gz"
	phpmyadmin_cmd_tar := exec.Command("curl", "-L", phpmyadmin_tar_url, "-o", "package.tar.gz")
	err := phpmyadmin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpmyadmin_zip_url := "https://files.phpmyadmin.net/phpMyAdmin/5.2.1/phpMyAdmin-5.2.1-all-languages.zip"
	phpmyadmin_cmd_zip := exec.Command("curl", "-L", phpmyadmin_zip_url, "-o", "package.zip")
	err = phpmyadmin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpmyadmin_bin_url := "https://files.phpmyadmin.net/phpMyAdmin/5.2.1/phpMyAdmin-5.2.1-all-languages.bin"
	phpmyadmin_cmd_bin := exec.Command("curl", "-L", phpmyadmin_bin_url, "-o", "binary.bin")
	err = phpmyadmin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpmyadmin_src_url := "https://files.phpmyadmin.net/phpMyAdmin/5.2.1/phpMyAdmin-5.2.1-all-languages.src.tar.gz"
	phpmyadmin_cmd_src := exec.Command("curl", "-L", phpmyadmin_src_url, "-o", "source.tar.gz")
	err = phpmyadmin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpmyadmin_cmd_direct := exec.Command("./binary")
	err = phpmyadmin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
