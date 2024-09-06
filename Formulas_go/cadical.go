package main

// Cadical - Clean and efficient state-of-the-art SAT solver
// Homepage: https://fmv.jku.at/cadical/

import (
	"fmt"
	
	"os/exec"
)

func installCadical() {
	// Método 1: Descargar y extraer .tar.gz
	cadical_tar_url := "https://github.com/arminbiere/cadical/archive/refs/tags/rel-2.0.0.tar.gz"
	cadical_cmd_tar := exec.Command("curl", "-L", cadical_tar_url, "-o", "package.tar.gz")
	err := cadical_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cadical_zip_url := "https://github.com/arminbiere/cadical/archive/refs/tags/rel-2.0.0.zip"
	cadical_cmd_zip := exec.Command("curl", "-L", cadical_zip_url, "-o", "package.zip")
	err = cadical_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cadical_bin_url := "https://github.com/arminbiere/cadical/archive/refs/tags/rel-2.0.0.bin"
	cadical_cmd_bin := exec.Command("curl", "-L", cadical_bin_url, "-o", "binary.bin")
	err = cadical_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cadical_src_url := "https://github.com/arminbiere/cadical/archive/refs/tags/rel-2.0.0.src.tar.gz"
	cadical_cmd_src := exec.Command("curl", "-L", cadical_src_url, "-o", "source.tar.gz")
	err = cadical_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cadical_cmd_direct := exec.Command("./binary")
	err = cadical_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
