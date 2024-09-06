package main

// Vilistextum - HTML to text converter
// Homepage: https://bhaak.net/vilistextum/

import (
	"fmt"
	
	"os/exec"
)

func installVilistextum() {
	// Método 1: Descargar y extraer .tar.gz
	vilistextum_tar_url := "https://bhaak.net/vilistextum/vilistextum-2.6.9.tar.gz"
	vilistextum_cmd_tar := exec.Command("curl", "-L", vilistextum_tar_url, "-o", "package.tar.gz")
	err := vilistextum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vilistextum_zip_url := "https://bhaak.net/vilistextum/vilistextum-2.6.9.zip"
	vilistextum_cmd_zip := exec.Command("curl", "-L", vilistextum_zip_url, "-o", "package.zip")
	err = vilistextum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vilistextum_bin_url := "https://bhaak.net/vilistextum/vilistextum-2.6.9.bin"
	vilistextum_cmd_bin := exec.Command("curl", "-L", vilistextum_bin_url, "-o", "binary.bin")
	err = vilistextum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vilistextum_src_url := "https://bhaak.net/vilistextum/vilistextum-2.6.9.src.tar.gz"
	vilistextum_cmd_src := exec.Command("curl", "-L", vilistextum_src_url, "-o", "source.tar.gz")
	err = vilistextum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vilistextum_cmd_direct := exec.Command("./binary")
	err = vilistextum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
