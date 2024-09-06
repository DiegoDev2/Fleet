package main

// Winetricks - Automatic workarounds for problems in Wine
// Homepage: https://github.com/Winetricks/winetricks

import (
	"fmt"
	
	"os/exec"
)

func installWinetricks() {
	// Método 1: Descargar y extraer .tar.gz
	winetricks_tar_url := "https://github.com/Winetricks/winetricks/archive/refs/tags/20240105.tar.gz"
	winetricks_cmd_tar := exec.Command("curl", "-L", winetricks_tar_url, "-o", "package.tar.gz")
	err := winetricks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	winetricks_zip_url := "https://github.com/Winetricks/winetricks/archive/refs/tags/20240105.zip"
	winetricks_cmd_zip := exec.Command("curl", "-L", winetricks_zip_url, "-o", "package.zip")
	err = winetricks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	winetricks_bin_url := "https://github.com/Winetricks/winetricks/archive/refs/tags/20240105.bin"
	winetricks_cmd_bin := exec.Command("curl", "-L", winetricks_bin_url, "-o", "binary.bin")
	err = winetricks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	winetricks_src_url := "https://github.com/Winetricks/winetricks/archive/refs/tags/20240105.src.tar.gz"
	winetricks_cmd_src := exec.Command("curl", "-L", winetricks_src_url, "-o", "source.tar.gz")
	err = winetricks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	winetricks_cmd_direct := exec.Command("./binary")
	err = winetricks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabextract")
	exec.Command("latte", "install", "cabextract").Run()
	fmt.Println("Instalando dependencia: p7zip")
	exec.Command("latte", "install", "p7zip").Run()
	fmt.Println("Instalando dependencia: unzip")
	exec.Command("latte", "install", "unzip").Run()
}
