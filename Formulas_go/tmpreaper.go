package main

// Tmpreaper - Clean up files in directories based on their age
// Homepage: https://packages.debian.org/sid/tmpreaper

import (
	"fmt"
	
	"os/exec"
)

func installTmpreaper() {
	// Método 1: Descargar y extraer .tar.gz
	tmpreaper_tar_url := "https://deb.debian.org/debian/pool/main/t/tmpreaper/tmpreaper_1.6.17.tar.gz"
	tmpreaper_cmd_tar := exec.Command("curl", "-L", tmpreaper_tar_url, "-o", "package.tar.gz")
	err := tmpreaper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmpreaper_zip_url := "https://deb.debian.org/debian/pool/main/t/tmpreaper/tmpreaper_1.6.17.zip"
	tmpreaper_cmd_zip := exec.Command("curl", "-L", tmpreaper_zip_url, "-o", "package.zip")
	err = tmpreaper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmpreaper_bin_url := "https://deb.debian.org/debian/pool/main/t/tmpreaper/tmpreaper_1.6.17.bin"
	tmpreaper_cmd_bin := exec.Command("curl", "-L", tmpreaper_bin_url, "-o", "binary.bin")
	err = tmpreaper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmpreaper_src_url := "https://deb.debian.org/debian/pool/main/t/tmpreaper/tmpreaper_1.6.17.src.tar.gz"
	tmpreaper_cmd_src := exec.Command("curl", "-L", tmpreaper_src_url, "-o", "source.tar.gz")
	err = tmpreaper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmpreaper_cmd_direct := exec.Command("./binary")
	err = tmpreaper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: e2fsprogs")
exec.Command("latte", "install", "e2fsprogs")
}
