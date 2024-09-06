package main

// Oscats - Computerized adaptive testing system
// Homepage: https://code.google.com/archive/p/oscats/

import (
	"fmt"
	
	"os/exec"
)

func installOscats() {
	// Método 1: Descargar y extraer .tar.gz
	oscats_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/oscats/oscats-0.6.tar.gz"
	oscats_cmd_tar := exec.Command("curl", "-L", oscats_tar_url, "-o", "package.tar.gz")
	err := oscats_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oscats_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/oscats/oscats-0.6.zip"
	oscats_cmd_zip := exec.Command("curl", "-L", oscats_zip_url, "-o", "package.zip")
	err = oscats_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oscats_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/oscats/oscats-0.6.bin"
	oscats_cmd_bin := exec.Command("curl", "-L", oscats_bin_url, "-o", "binary.bin")
	err = oscats_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oscats_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/oscats/oscats-0.6.src.tar.gz"
	oscats_cmd_src := exec.Command("curl", "-L", oscats_src_url, "-o", "source.tar.gz")
	err = oscats_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oscats_cmd_direct := exec.Command("./binary")
	err = oscats_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
}
