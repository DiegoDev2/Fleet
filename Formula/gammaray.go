package main

// Gammaray - Examine and manipulate Qt application internals at runtime
// Homepage: https://www.kdab.com/gammaray

import (
	"fmt"
	
	"os/exec"
)

func installGammaray() {
	// Método 1: Descargar y extraer .tar.gz
	gammaray_tar_url := "https://github.com/KDAB/GammaRay/releases/download/v3.1.0/gammaray-3.1.0.tar.gz"
	gammaray_cmd_tar := exec.Command("curl", "-L", gammaray_tar_url, "-o", "package.tar.gz")
	err := gammaray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gammaray_zip_url := "https://github.com/KDAB/GammaRay/releases/download/v3.1.0/gammaray-3.1.0.zip"
	gammaray_cmd_zip := exec.Command("curl", "-L", gammaray_zip_url, "-o", "package.zip")
	err = gammaray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gammaray_bin_url := "https://github.com/KDAB/GammaRay/releases/download/v3.1.0/gammaray-3.1.0.bin"
	gammaray_cmd_bin := exec.Command("curl", "-L", gammaray_bin_url, "-o", "binary.bin")
	err = gammaray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gammaray_src_url := "https://github.com/KDAB/GammaRay/releases/download/v3.1.0/gammaray-3.1.0.src.tar.gz"
	gammaray_cmd_src := exec.Command("curl", "-L", gammaray_src_url, "-o", "source.tar.gz")
	err = gammaray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gammaray_cmd_direct := exec.Command("./binary")
	err = gammaray_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
