package main

// PariSeadataBig - Additional modular polynomial data for PARI/GP
// Homepage: https://pari.math.u-bordeaux.fr/packages.html

import (
	"fmt"
	
	"os/exec"
)

func installPariSeadataBig() {
	// Método 1: Descargar y extraer .tar.gz
	pariseadatabig_tar_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata-big.tar"
	pariseadatabig_cmd_tar := exec.Command("curl", "-L", pariseadatabig_tar_url, "-o", "package.tar.gz")
	err := pariseadatabig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pariseadatabig_zip_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata-big.tar"
	pariseadatabig_cmd_zip := exec.Command("curl", "-L", pariseadatabig_zip_url, "-o", "package.zip")
	err = pariseadatabig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pariseadatabig_bin_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata-big.tar"
	pariseadatabig_cmd_bin := exec.Command("curl", "-L", pariseadatabig_bin_url, "-o", "binary.bin")
	err = pariseadatabig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pariseadatabig_src_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata-big.tar"
	pariseadatabig_cmd_src := exec.Command("curl", "-L", pariseadatabig_src_url, "-o", "source.tar.gz")
	err = pariseadatabig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pariseadatabig_cmd_direct := exec.Command("./binary")
	err = pariseadatabig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pari")
exec.Command("latte", "install", "pari")
	fmt.Println("Instalando dependencia: pari-seadata")
exec.Command("latte", "install", "pari-seadata")
}
