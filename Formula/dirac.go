package main

// Dirac - General-purpose video codec aimed at a range of resolutions
// Homepage: https://sourceforge.net/projects/dirac/

import (
	"fmt"
	
	"os/exec"
)

func installDirac() {
	// Método 1: Descargar y extraer .tar.gz
	dirac_tar_url := "https://downloads.sourceforge.net/project/dirac/dirac-codec/Dirac-1.0.2/dirac-1.0.2.tar.gz"
	dirac_cmd_tar := exec.Command("curl", "-L", dirac_tar_url, "-o", "package.tar.gz")
	err := dirac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dirac_zip_url := "https://downloads.sourceforge.net/project/dirac/dirac-codec/Dirac-1.0.2/dirac-1.0.2.zip"
	dirac_cmd_zip := exec.Command("curl", "-L", dirac_zip_url, "-o", "package.zip")
	err = dirac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dirac_bin_url := "https://downloads.sourceforge.net/project/dirac/dirac-codec/Dirac-1.0.2/dirac-1.0.2.bin"
	dirac_cmd_bin := exec.Command("curl", "-L", dirac_bin_url, "-o", "binary.bin")
	err = dirac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dirac_src_url := "https://downloads.sourceforge.net/project/dirac/dirac-codec/Dirac-1.0.2/dirac-1.0.2.src.tar.gz"
	dirac_cmd_src := exec.Command("curl", "-L", dirac_src_url, "-o", "source.tar.gz")
	err = dirac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dirac_cmd_direct := exec.Command("./binary")
	err = dirac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
