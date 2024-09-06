package main

// Ancient - Decompression routines for ancient formats
// Homepage: https://github.com/temisu/ancient

import (
	"fmt"
	
	"os/exec"
)

func installAncient() {
	// Método 1: Descargar y extraer .tar.gz
	ancient_tar_url := "https://github.com/temisu/ancient/archive/refs/tags/v2.2.0.tar.gz"
	ancient_cmd_tar := exec.Command("curl", "-L", ancient_tar_url, "-o", "package.tar.gz")
	err := ancient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ancient_zip_url := "https://github.com/temisu/ancient/archive/refs/tags/v2.2.0.zip"
	ancient_cmd_zip := exec.Command("curl", "-L", ancient_zip_url, "-o", "package.zip")
	err = ancient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ancient_bin_url := "https://github.com/temisu/ancient/archive/refs/tags/v2.2.0.bin"
	ancient_cmd_bin := exec.Command("curl", "-L", ancient_bin_url, "-o", "binary.bin")
	err = ancient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ancient_src_url := "https://github.com/temisu/ancient/archive/refs/tags/v2.2.0.src.tar.gz"
	ancient_cmd_src := exec.Command("curl", "-L", ancient_src_url, "-o", "source.tar.gz")
	err = ancient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ancient_cmd_direct := exec.Command("./binary")
	err = ancient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: autoconf-archive")
	exec.Command("latte", "install", "autoconf-archive").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
