package main

// Qprint - Encoder and decoder for quoted-printable encoding
// Homepage: https://www.fourmilab.ch/webtools/qprint/

import (
	"fmt"
	
	"os/exec"
)

func installQprint() {
	// Método 1: Descargar y extraer .tar.gz
	qprint_tar_url := "https://www.fourmilab.ch/webtools/qprint/qprint-1.1.tar.gz"
	qprint_cmd_tar := exec.Command("curl", "-L", qprint_tar_url, "-o", "package.tar.gz")
	err := qprint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qprint_zip_url := "https://www.fourmilab.ch/webtools/qprint/qprint-1.1.zip"
	qprint_cmd_zip := exec.Command("curl", "-L", qprint_zip_url, "-o", "package.zip")
	err = qprint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qprint_bin_url := "https://www.fourmilab.ch/webtools/qprint/qprint-1.1.bin"
	qprint_cmd_bin := exec.Command("curl", "-L", qprint_bin_url, "-o", "binary.bin")
	err = qprint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qprint_src_url := "https://www.fourmilab.ch/webtools/qprint/qprint-1.1.src.tar.gz"
	qprint_cmd_src := exec.Command("curl", "-L", qprint_src_url, "-o", "source.tar.gz")
	err = qprint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qprint_cmd_direct := exec.Command("./binary")
	err = qprint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
