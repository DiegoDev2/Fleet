package main

// Slashem - Fork/variant of Nethack
// Homepage: https://slashem.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSlashem() {
	// Método 1: Descargar y extraer .tar.gz
	slashem_tar_url := "https://downloads.sourceforge.net/project/slashem/slashem-source/0.0.8E0F1/se008e0f1.tar.gz"
	slashem_cmd_tar := exec.Command("curl", "-L", slashem_tar_url, "-o", "package.tar.gz")
	err := slashem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slashem_zip_url := "https://downloads.sourceforge.net/project/slashem/slashem-source/0.0.8E0F1/se008e0f1.zip"
	slashem_cmd_zip := exec.Command("curl", "-L", slashem_zip_url, "-o", "package.zip")
	err = slashem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slashem_bin_url := "https://downloads.sourceforge.net/project/slashem/slashem-source/0.0.8E0F1/se008e0f1.bin"
	slashem_cmd_bin := exec.Command("curl", "-L", slashem_bin_url, "-o", "binary.bin")
	err = slashem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slashem_src_url := "https://downloads.sourceforge.net/project/slashem/slashem-source/0.0.8E0F1/se008e0f1.src.tar.gz"
	slashem_cmd_src := exec.Command("curl", "-L", slashem_src_url, "-o", "source.tar.gz")
	err = slashem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slashem_cmd_direct := exec.Command("./binary")
	err = slashem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
