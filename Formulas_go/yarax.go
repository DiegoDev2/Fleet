package main

// YaraX - Tool to do pattern matching for malware research
// Homepage: https://virustotal.github.io/yara-x/

import (
	"fmt"
	
	"os/exec"
)

func installYaraX() {
	// Método 1: Descargar y extraer .tar.gz
	yarax_tar_url := "https://github.com/VirusTotal/yara-x/archive/refs/tags/v0.7.0.tar.gz"
	yarax_cmd_tar := exec.Command("curl", "-L", yarax_tar_url, "-o", "package.tar.gz")
	err := yarax_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yarax_zip_url := "https://github.com/VirusTotal/yara-x/archive/refs/tags/v0.7.0.zip"
	yarax_cmd_zip := exec.Command("curl", "-L", yarax_zip_url, "-o", "package.zip")
	err = yarax_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yarax_bin_url := "https://github.com/VirusTotal/yara-x/archive/refs/tags/v0.7.0.bin"
	yarax_cmd_bin := exec.Command("curl", "-L", yarax_bin_url, "-o", "binary.bin")
	err = yarax_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yarax_src_url := "https://github.com/VirusTotal/yara-x/archive/refs/tags/v0.7.0.src.tar.gz"
	yarax_cmd_src := exec.Command("curl", "-L", yarax_src_url, "-o", "source.tar.gz")
	err = yarax_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yarax_cmd_direct := exec.Command("./binary")
	err = yarax_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
