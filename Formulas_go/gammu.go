package main

// Gammu - Command-line utility to control a phone
// Homepage: https://wammu.eu/gammu/

import (
	"fmt"
	
	"os/exec"
)

func installGammu() {
	// Método 1: Descargar y extraer .tar.gz
	gammu_tar_url := "https://dl.cihar.com/gammu/releases/gammu-1.42.0.tar.xz"
	gammu_cmd_tar := exec.Command("curl", "-L", gammu_tar_url, "-o", "package.tar.gz")
	err := gammu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gammu_zip_url := "https://dl.cihar.com/gammu/releases/gammu-1.42.0.tar.xz"
	gammu_cmd_zip := exec.Command("curl", "-L", gammu_zip_url, "-o", "package.zip")
	err = gammu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gammu_bin_url := "https://dl.cihar.com/gammu/releases/gammu-1.42.0.tar.xz"
	gammu_cmd_bin := exec.Command("curl", "-L", gammu_bin_url, "-o", "binary.bin")
	err = gammu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gammu_src_url := "https://dl.cihar.com/gammu/releases/gammu-1.42.0.tar.xz"
	gammu_cmd_src := exec.Command("curl", "-L", gammu_src_url, "-o", "source.tar.gz")
	err = gammu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gammu_cmd_direct := exec.Command("./binary")
	err = gammu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
