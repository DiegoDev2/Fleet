package main

// Muon - Meson-compatible build system
// Homepage: https://muon.build

import (
	"fmt"
	
	"os/exec"
)

func installMuon() {
	// Método 1: Descargar y extraer .tar.gz
	muon_tar_url := "https://git.sr.ht/~lattis/muon/archive/0.2.0.tar.gz"
	muon_cmd_tar := exec.Command("curl", "-L", muon_tar_url, "-o", "package.tar.gz")
	err := muon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	muon_zip_url := "https://git.sr.ht/~lattis/muon/archive/0.2.0.zip"
	muon_cmd_zip := exec.Command("curl", "-L", muon_zip_url, "-o", "package.zip")
	err = muon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	muon_bin_url := "https://git.sr.ht/~lattis/muon/archive/0.2.0.bin"
	muon_cmd_bin := exec.Command("curl", "-L", muon_bin_url, "-o", "binary.bin")
	err = muon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	muon_src_url := "https://git.sr.ht/~lattis/muon/archive/0.2.0.src.tar.gz"
	muon_cmd_src := exec.Command("curl", "-L", muon_src_url, "-o", "source.tar.gz")
	err = muon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	muon_cmd_direct := exec.Command("./binary")
	err = muon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
