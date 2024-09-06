package main

// Slrn - Powerful console-based newsreader
// Homepage: https://slrn.info/

import (
	"fmt"
	
	"os/exec"
)

func installSlrn() {
	// Método 1: Descargar y extraer .tar.gz
	slrn_tar_url := "https://jedsoft.org/releases/slrn/slrn-1.0.3a.tar.bz2"
	slrn_cmd_tar := exec.Command("curl", "-L", slrn_tar_url, "-o", "package.tar.gz")
	err := slrn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slrn_zip_url := "https://jedsoft.org/releases/slrn/slrn-1.0.3a.tar.bz2"
	slrn_cmd_zip := exec.Command("curl", "-L", slrn_zip_url, "-o", "package.zip")
	err = slrn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slrn_bin_url := "https://jedsoft.org/releases/slrn/slrn-1.0.3a.tar.bz2"
	slrn_cmd_bin := exec.Command("curl", "-L", slrn_bin_url, "-o", "binary.bin")
	err = slrn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slrn_src_url := "https://jedsoft.org/releases/slrn/slrn-1.0.3a.tar.bz2"
	slrn_cmd_src := exec.Command("curl", "-L", slrn_src_url, "-o", "source.tar.gz")
	err = slrn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slrn_cmd_direct := exec.Command("./binary")
	err = slrn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: s-lang")
exec.Command("latte", "install", "s-lang")
}
