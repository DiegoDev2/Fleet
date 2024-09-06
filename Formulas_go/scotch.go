package main

// Scotch - Package for graph partitioning, graph clustering, and sparse matrix ordering
// Homepage: https://gitlab.inria.fr/scotch/scotch

import (
	"fmt"
	
	"os/exec"
)

func installScotch() {
	// Método 1: Descargar y extraer .tar.gz
	scotch_tar_url := "https://gitlab.inria.fr/scotch/scotch/-/archive/v7.0.5/scotch-v7.0.5.tar.bz2"
	scotch_cmd_tar := exec.Command("curl", "-L", scotch_tar_url, "-o", "package.tar.gz")
	err := scotch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scotch_zip_url := "https://gitlab.inria.fr/scotch/scotch/-/archive/v7.0.5/scotch-v7.0.5.tar.bz2"
	scotch_cmd_zip := exec.Command("curl", "-L", scotch_zip_url, "-o", "package.zip")
	err = scotch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scotch_bin_url := "https://gitlab.inria.fr/scotch/scotch/-/archive/v7.0.5/scotch-v7.0.5.tar.bz2"
	scotch_cmd_bin := exec.Command("curl", "-L", scotch_bin_url, "-o", "binary.bin")
	err = scotch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scotch_src_url := "https://gitlab.inria.fr/scotch/scotch/-/archive/v7.0.5/scotch-v7.0.5.tar.bz2"
	scotch_cmd_src := exec.Command("curl", "-L", scotch_src_url, "-o", "source.tar.gz")
	err = scotch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scotch_cmd_direct := exec.Command("./binary")
	err = scotch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
}
