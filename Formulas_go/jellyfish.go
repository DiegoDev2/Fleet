package main

// Jellyfish - Fast, memory-efficient counting of DNA k-mers
// Homepage: https://github.com/gmarcais/Jellyfish

import (
	"fmt"
	
	"os/exec"
)

func installJellyfish() {
	// Método 1: Descargar y extraer .tar.gz
	jellyfish_tar_url := "https://github.com/gmarcais/Jellyfish/releases/download/v2.3.1/jellyfish-2.3.1.tar.gz"
	jellyfish_cmd_tar := exec.Command("curl", "-L", jellyfish_tar_url, "-o", "package.tar.gz")
	err := jellyfish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jellyfish_zip_url := "https://github.com/gmarcais/Jellyfish/releases/download/v2.3.1/jellyfish-2.3.1.zip"
	jellyfish_cmd_zip := exec.Command("curl", "-L", jellyfish_zip_url, "-o", "package.zip")
	err = jellyfish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jellyfish_bin_url := "https://github.com/gmarcais/Jellyfish/releases/download/v2.3.1/jellyfish-2.3.1.bin"
	jellyfish_cmd_bin := exec.Command("curl", "-L", jellyfish_bin_url, "-o", "binary.bin")
	err = jellyfish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jellyfish_src_url := "https://github.com/gmarcais/Jellyfish/releases/download/v2.3.1/jellyfish-2.3.1.src.tar.gz"
	jellyfish_cmd_src := exec.Command("curl", "-L", jellyfish_src_url, "-o", "source.tar.gz")
	err = jellyfish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jellyfish_cmd_direct := exec.Command("./binary")
	err = jellyfish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: htslib")
exec.Command("latte", "install", "htslib")
}
