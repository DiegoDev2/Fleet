package main

// Hypre - Library featuring parallel multigrid methods for grid problems
// Homepage: https://computing.llnl.gov/projects/hypre-scalable-linear-solvers-multigrid-methods

import (
	"fmt"
	
	"os/exec"
)

func installHypre() {
	// Método 1: Descargar y extraer .tar.gz
	hypre_tar_url := "https://github.com/hypre-space/hypre/archive/refs/tags/v2.31.0.tar.gz"
	hypre_cmd_tar := exec.Command("curl", "-L", hypre_tar_url, "-o", "package.tar.gz")
	err := hypre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hypre_zip_url := "https://github.com/hypre-space/hypre/archive/refs/tags/v2.31.0.zip"
	hypre_cmd_zip := exec.Command("curl", "-L", hypre_zip_url, "-o", "package.zip")
	err = hypre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hypre_bin_url := "https://github.com/hypre-space/hypre/archive/refs/tags/v2.31.0.bin"
	hypre_cmd_bin := exec.Command("curl", "-L", hypre_bin_url, "-o", "binary.bin")
	err = hypre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hypre_src_url := "https://github.com/hypre-space/hypre/archive/refs/tags/v2.31.0.src.tar.gz"
	hypre_cmd_src := exec.Command("curl", "-L", hypre_src_url, "-o", "source.tar.gz")
	err = hypre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hypre_cmd_direct := exec.Command("./binary")
	err = hypre_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
}
