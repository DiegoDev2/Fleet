package main

// Smap - Drop-in replacement for Nmap powered by shodan.io
// Homepage: https://github.com/s0md3v/Smap

import (
	"fmt"
	
	"os/exec"
)

func installSmap() {
	// Método 1: Descargar y extraer .tar.gz
	smap_tar_url := "https://github.com/s0md3v/Smap/archive/refs/tags/0.1.12.tar.gz"
	smap_cmd_tar := exec.Command("curl", "-L", smap_tar_url, "-o", "package.tar.gz")
	err := smap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smap_zip_url := "https://github.com/s0md3v/Smap/archive/refs/tags/0.1.12.zip"
	smap_cmd_zip := exec.Command("curl", "-L", smap_zip_url, "-o", "package.zip")
	err = smap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smap_bin_url := "https://github.com/s0md3v/Smap/archive/refs/tags/0.1.12.bin"
	smap_cmd_bin := exec.Command("curl", "-L", smap_bin_url, "-o", "binary.bin")
	err = smap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smap_src_url := "https://github.com/s0md3v/Smap/archive/refs/tags/0.1.12.src.tar.gz"
	smap_cmd_src := exec.Command("curl", "-L", smap_src_url, "-o", "source.tar.gz")
	err = smap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smap_cmd_direct := exec.Command("./binary")
	err = smap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
