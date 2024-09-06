package main

// S6 - Small & secure supervision software suite
// Homepage: https://skarnet.org/software/s6/

import (
	"fmt"
	
	"os/exec"
)

func installS6() {
	// Método 1: Descargar y extraer .tar.gz
	s6_tar_url := "https://skarnet.org/software/s6/s6-2.13.0.0.tar.gz"
	s6_cmd_tar := exec.Command("curl", "-L", s6_tar_url, "-o", "package.tar.gz")
	err := s6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s6_zip_url := "https://skarnet.org/software/s6/s6-2.13.0.0.zip"
	s6_cmd_zip := exec.Command("curl", "-L", s6_zip_url, "-o", "package.zip")
	err = s6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s6_bin_url := "https://skarnet.org/software/s6/s6-2.13.0.0.bin"
	s6_cmd_bin := exec.Command("curl", "-L", s6_bin_url, "-o", "binary.bin")
	err = s6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s6_src_url := "https://skarnet.org/software/s6/s6-2.13.0.0.src.tar.gz"
	s6_cmd_src := exec.Command("curl", "-L", s6_src_url, "-o", "source.tar.gz")
	err = s6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s6_cmd_direct := exec.Command("./binary")
	err = s6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
