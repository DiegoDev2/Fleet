package main

// StoneSoup - Dungeon Crawl Stone Soup: a roguelike game
// Homepage: https://crawl.develz.org/

import (
	"fmt"
	
	"os/exec"
)

func installStoneSoup() {
	// Método 1: Descargar y extraer .tar.gz
	stonesoup_tar_url := "https://github.com/crawl/crawl/archive/refs/tags/0.30.1.tar.gz"
	stonesoup_cmd_tar := exec.Command("curl", "-L", stonesoup_tar_url, "-o", "package.tar.gz")
	err := stonesoup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stonesoup_zip_url := "https://github.com/crawl/crawl/archive/refs/tags/0.30.1.zip"
	stonesoup_cmd_zip := exec.Command("curl", "-L", stonesoup_zip_url, "-o", "package.zip")
	err = stonesoup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stonesoup_bin_url := "https://github.com/crawl/crawl/archive/refs/tags/0.30.1.bin"
	stonesoup_cmd_bin := exec.Command("curl", "-L", stonesoup_bin_url, "-o", "binary.bin")
	err = stonesoup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stonesoup_src_url := "https://github.com/crawl/crawl/archive/refs/tags/0.30.1.src.tar.gz"
	stonesoup_cmd_src := exec.Command("curl", "-L", stonesoup_src_url, "-o", "source.tar.gz")
	err = stonesoup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stonesoup_cmd_direct := exec.Command("./binary")
	err = stonesoup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: pyyaml")
	exec.Command("latte", "install", "pyyaml").Run()
	fmt.Println("Instalando dependencia: lua@5.1")
	exec.Command("latte", "install", "lua@5.1").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
