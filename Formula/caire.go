package main

// Caire - Content aware image resize tool
// Homepage: https://github.com/esimov/caire

import (
	"fmt"
	
	"os/exec"
)

func installCaire() {
	// Método 1: Descargar y extraer .tar.gz
	caire_tar_url := "https://github.com/esimov/caire/archive/refs/tags/v1.4.6.tar.gz"
	caire_cmd_tar := exec.Command("curl", "-L", caire_tar_url, "-o", "package.tar.gz")
	err := caire_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	caire_zip_url := "https://github.com/esimov/caire/archive/refs/tags/v1.4.6.zip"
	caire_cmd_zip := exec.Command("curl", "-L", caire_zip_url, "-o", "package.zip")
	err = caire_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	caire_bin_url := "https://github.com/esimov/caire/archive/refs/tags/v1.4.6.bin"
	caire_cmd_bin := exec.Command("curl", "-L", caire_bin_url, "-o", "binary.bin")
	err = caire_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	caire_src_url := "https://github.com/esimov/caire/archive/refs/tags/v1.4.6.src.tar.gz"
	caire_cmd_src := exec.Command("curl", "-L", caire_src_url, "-o", "source.tar.gz")
	err = caire_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	caire_cmd_direct := exec.Command("./binary")
	err = caire_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vulkan-headers")
	exec.Command("latte", "install", "vulkan-headers").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: libxfixes")
	exec.Command("latte", "install", "libxfixes").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
