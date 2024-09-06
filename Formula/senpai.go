package main

// Senpai - Modern terminal IRC client
// Homepage: https://sr.ht/~delthas/senpai/

import (
	"fmt"
	
	"os/exec"
)

func installSenpai() {
	// Método 1: Descargar y extraer .tar.gz
	senpai_tar_url := "https://git.sr.ht/~delthas/senpai/archive/v0.3.0.tar.gz"
	senpai_cmd_tar := exec.Command("curl", "-L", senpai_tar_url, "-o", "package.tar.gz")
	err := senpai_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	senpai_zip_url := "https://git.sr.ht/~delthas/senpai/archive/v0.3.0.zip"
	senpai_cmd_zip := exec.Command("curl", "-L", senpai_zip_url, "-o", "package.zip")
	err = senpai_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	senpai_bin_url := "https://git.sr.ht/~delthas/senpai/archive/v0.3.0.bin"
	senpai_cmd_bin := exec.Command("curl", "-L", senpai_bin_url, "-o", "binary.bin")
	err = senpai_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	senpai_src_url := "https://git.sr.ht/~delthas/senpai/archive/v0.3.0.src.tar.gz"
	senpai_cmd_src := exec.Command("curl", "-L", senpai_src_url, "-o", "source.tar.gz")
	err = senpai_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	senpai_cmd_direct := exec.Command("./binary")
	err = senpai_cmd_direct.Run()
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
