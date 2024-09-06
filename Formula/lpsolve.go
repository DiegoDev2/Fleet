package main

// LpSolve - Mixed integer linear programming solver
// Homepage: https://sourceforge.net/projects/lpsolve/

import (
	"fmt"
	
	"os/exec"
)

func installLpSolve() {
	// Método 1: Descargar y extraer .tar.gz
	lpsolve_tar_url := "https://downloads.sourceforge.net/lpsolve/lp_solve_5.5.2.11_source.tar.gz"
	lpsolve_cmd_tar := exec.Command("curl", "-L", lpsolve_tar_url, "-o", "package.tar.gz")
	err := lpsolve_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lpsolve_zip_url := "https://downloads.sourceforge.net/lpsolve/lp_solve_5.5.2.11_source.zip"
	lpsolve_cmd_zip := exec.Command("curl", "-L", lpsolve_zip_url, "-o", "package.zip")
	err = lpsolve_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lpsolve_bin_url := "https://downloads.sourceforge.net/lpsolve/lp_solve_5.5.2.11_source.bin"
	lpsolve_cmd_bin := exec.Command("curl", "-L", lpsolve_bin_url, "-o", "binary.bin")
	err = lpsolve_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lpsolve_src_url := "https://downloads.sourceforge.net/lpsolve/lp_solve_5.5.2.11_source.src.tar.gz"
	lpsolve_cmd_src := exec.Command("curl", "-L", lpsolve_src_url, "-o", "source.tar.gz")
	err = lpsolve_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lpsolve_cmd_direct := exec.Command("./binary")
	err = lpsolve_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
