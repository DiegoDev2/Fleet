package main

// Ltl2ba - Translate LTL formulae to Buchi automata
// Homepage: https://www.lsv.ens-cachan.fr/~gastin/ltl2ba/

import (
	"fmt"
	
	"os/exec"
)

func installLtl2ba() {
	// Método 1: Descargar y extraer .tar.gz
	ltl2ba_tar_url := "https://www.lsv.fr/~gastin/ltl2ba/ltl2ba-1.3.tar.gz"
	ltl2ba_cmd_tar := exec.Command("curl", "-L", ltl2ba_tar_url, "-o", "package.tar.gz")
	err := ltl2ba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ltl2ba_zip_url := "https://www.lsv.fr/~gastin/ltl2ba/ltl2ba-1.3.zip"
	ltl2ba_cmd_zip := exec.Command("curl", "-L", ltl2ba_zip_url, "-o", "package.zip")
	err = ltl2ba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ltl2ba_bin_url := "https://www.lsv.fr/~gastin/ltl2ba/ltl2ba-1.3.bin"
	ltl2ba_cmd_bin := exec.Command("curl", "-L", ltl2ba_bin_url, "-o", "binary.bin")
	err = ltl2ba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ltl2ba_src_url := "https://www.lsv.fr/~gastin/ltl2ba/ltl2ba-1.3.src.tar.gz"
	ltl2ba_cmd_src := exec.Command("curl", "-L", ltl2ba_src_url, "-o", "source.tar.gz")
	err = ltl2ba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ltl2ba_cmd_direct := exec.Command("./binary")
	err = ltl2ba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
