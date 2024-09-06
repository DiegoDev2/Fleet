package main

// Gif2png - Convert GIFs to PNGs
// Homepage: http://www.catb.org/~esr/gif2png/

import (
	"fmt"
	
	"os/exec"
)

func installGif2png() {
	// Método 1: Descargar y extraer .tar.gz
	gif2png_tar_url := "https://gitlab.com/esr/gif2png/-/archive/3.0.3/gif2png-3.0.3.tar.bz2"
	gif2png_cmd_tar := exec.Command("curl", "-L", gif2png_tar_url, "-o", "package.tar.gz")
	err := gif2png_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gif2png_zip_url := "https://gitlab.com/esr/gif2png/-/archive/3.0.3/gif2png-3.0.3.tar.bz2"
	gif2png_cmd_zip := exec.Command("curl", "-L", gif2png_zip_url, "-o", "package.zip")
	err = gif2png_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gif2png_bin_url := "https://gitlab.com/esr/gif2png/-/archive/3.0.3/gif2png-3.0.3.tar.bz2"
	gif2png_cmd_bin := exec.Command("curl", "-L", gif2png_bin_url, "-o", "binary.bin")
	err = gif2png_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gif2png_src_url := "https://gitlab.com/esr/gif2png/-/archive/3.0.3/gif2png-3.0.3.tar.bz2"
	gif2png_cmd_src := exec.Command("curl", "-L", gif2png_src_url, "-o", "source.tar.gz")
	err = gif2png_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gif2png_cmd_direct := exec.Command("./binary")
	err = gif2png_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
}
