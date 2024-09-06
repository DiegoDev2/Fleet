package main

// Hut - CLI tool for sr.ht
// Homepage: https://sr.ht/~xenrox/hut

import (
	"fmt"
	
	"os/exec"
)

func installHut() {
	// Método 1: Descargar y extraer .tar.gz
	hut_tar_url := "https://git.sr.ht/~xenrox/hut/archive/v0.6.0.tar.gz"
	hut_cmd_tar := exec.Command("curl", "-L", hut_tar_url, "-o", "package.tar.gz")
	err := hut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hut_zip_url := "https://git.sr.ht/~xenrox/hut/archive/v0.6.0.zip"
	hut_cmd_zip := exec.Command("curl", "-L", hut_zip_url, "-o", "package.zip")
	err = hut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hut_bin_url := "https://git.sr.ht/~xenrox/hut/archive/v0.6.0.bin"
	hut_cmd_bin := exec.Command("curl", "-L", hut_bin_url, "-o", "binary.bin")
	err = hut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hut_src_url := "https://git.sr.ht/~xenrox/hut/archive/v0.6.0.src.tar.gz"
	hut_cmd_src := exec.Command("curl", "-L", hut_src_url, "-o", "source.tar.gz")
	err = hut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hut_cmd_direct := exec.Command("./binary")
	err = hut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
}
