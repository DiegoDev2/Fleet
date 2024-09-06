package main

// Platypus - Create macOS applications from {Perl,Ruby,sh,Python} scripts
// Homepage: https://sveinbjorn.org/platypus

import (
	"fmt"
	
	"os/exec"
)

func installPlatypus() {
	// Método 1: Descargar y extraer .tar.gz
	platypus_tar_url := "https://sveinbjorn.org/files/software/platypus/platypus5.3.src.zip"
	platypus_cmd_tar := exec.Command("curl", "-L", platypus_tar_url, "-o", "package.tar.gz")
	err := platypus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	platypus_zip_url := "https://sveinbjorn.org/files/software/platypus/platypus5.3.src.zip"
	platypus_cmd_zip := exec.Command("curl", "-L", platypus_zip_url, "-o", "package.zip")
	err = platypus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	platypus_bin_url := "https://sveinbjorn.org/files/software/platypus/platypus5.3.src.zip"
	platypus_cmd_bin := exec.Command("curl", "-L", platypus_bin_url, "-o", "binary.bin")
	err = platypus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	platypus_src_url := "https://sveinbjorn.org/files/software/platypus/platypus5.3.src.zip"
	platypus_cmd_src := exec.Command("curl", "-L", platypus_src_url, "-o", "source.tar.gz")
	err = platypus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	platypus_cmd_direct := exec.Command("./binary")
	err = platypus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
