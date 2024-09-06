package main

// Cfonts - Sexy ANSI fonts for the console
// Homepage: https://github.com/dominikwilkowski/cfonts

import (
	"fmt"
	
	"os/exec"
)

func installCfonts() {
	// Método 1: Descargar y extraer .tar.gz
	cfonts_tar_url := "https://github.com/dominikwilkowski/cfonts/archive/refs/tags/v1.2.0rust.tar.gz"
	cfonts_cmd_tar := exec.Command("curl", "-L", cfonts_tar_url, "-o", "package.tar.gz")
	err := cfonts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfonts_zip_url := "https://github.com/dominikwilkowski/cfonts/archive/refs/tags/v1.2.0rust.zip"
	cfonts_cmd_zip := exec.Command("curl", "-L", cfonts_zip_url, "-o", "package.zip")
	err = cfonts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfonts_bin_url := "https://github.com/dominikwilkowski/cfonts/archive/refs/tags/v1.2.0rust.bin"
	cfonts_cmd_bin := exec.Command("curl", "-L", cfonts_bin_url, "-o", "binary.bin")
	err = cfonts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfonts_src_url := "https://github.com/dominikwilkowski/cfonts/archive/refs/tags/v1.2.0rust.src.tar.gz"
	cfonts_cmd_src := exec.Command("curl", "-L", cfonts_src_url, "-o", "source.tar.gz")
	err = cfonts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfonts_cmd_direct := exec.Command("./binary")
	err = cfonts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
