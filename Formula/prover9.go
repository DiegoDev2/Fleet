package main

// Prover9 - Automated theorem prover for first-order and equational logic
// Homepage: https://www.cs.unm.edu/~mccune/prover9/

import (
	"fmt"
	
	"os/exec"
)

func installProver9() {
	// Método 1: Descargar y extraer .tar.gz
	prover9_tar_url := "https://www.cs.unm.edu/~mccune/prover9/download/LADR-2009-11A.tar.gz"
	prover9_cmd_tar := exec.Command("curl", "-L", prover9_tar_url, "-o", "package.tar.gz")
	err := prover9_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prover9_zip_url := "https://www.cs.unm.edu/~mccune/prover9/download/LADR-2009-11A.zip"
	prover9_cmd_zip := exec.Command("curl", "-L", prover9_zip_url, "-o", "package.zip")
	err = prover9_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prover9_bin_url := "https://www.cs.unm.edu/~mccune/prover9/download/LADR-2009-11A.bin"
	prover9_cmd_bin := exec.Command("curl", "-L", prover9_bin_url, "-o", "binary.bin")
	err = prover9_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prover9_src_url := "https://www.cs.unm.edu/~mccune/prover9/download/LADR-2009-11A.src.tar.gz"
	prover9_cmd_src := exec.Command("curl", "-L", prover9_src_url, "-o", "source.tar.gz")
	err = prover9_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prover9_cmd_direct := exec.Command("./binary")
	err = prover9_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
