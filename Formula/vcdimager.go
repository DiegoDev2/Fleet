package main

// Vcdimager - (Super) video CD authoring solution
// Homepage: https://www.gnu.org/software/vcdimager/

import (
	"fmt"
	
	"os/exec"
)

func installVcdimager() {
	// Método 1: Descargar y extraer .tar.gz
	vcdimager_tar_url := "https://ftp.gnu.org/gnu/vcdimager/vcdimager-2.0.1.tar.gz"
	vcdimager_cmd_tar := exec.Command("curl", "-L", vcdimager_tar_url, "-o", "package.tar.gz")
	err := vcdimager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vcdimager_zip_url := "https://ftp.gnu.org/gnu/vcdimager/vcdimager-2.0.1.zip"
	vcdimager_cmd_zip := exec.Command("curl", "-L", vcdimager_zip_url, "-o", "package.zip")
	err = vcdimager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vcdimager_bin_url := "https://ftp.gnu.org/gnu/vcdimager/vcdimager-2.0.1.bin"
	vcdimager_cmd_bin := exec.Command("curl", "-L", vcdimager_bin_url, "-o", "binary.bin")
	err = vcdimager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vcdimager_src_url := "https://ftp.gnu.org/gnu/vcdimager/vcdimager-2.0.1.src.tar.gz"
	vcdimager_cmd_src := exec.Command("curl", "-L", vcdimager_src_url, "-o", "source.tar.gz")
	err = vcdimager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vcdimager_cmd_direct := exec.Command("./binary")
	err = vcdimager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libcdio")
	exec.Command("latte", "install", "libcdio").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
}
