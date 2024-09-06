package main

// PariSeadata - Modular polynomial data for PARI/GP
// Homepage: https://pari.math.u-bordeaux.fr/packages.html

import (
	"fmt"
	
	"os/exec"
)

func installPariSeadata() {
	// Método 1: Descargar y extraer .tar.gz
	pariseadata_tar_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata.tgz"
	pariseadata_cmd_tar := exec.Command("curl", "-L", pariseadata_tar_url, "-o", "package.tar.gz")
	err := pariseadata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pariseadata_zip_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata.tgz"
	pariseadata_cmd_zip := exec.Command("curl", "-L", pariseadata_zip_url, "-o", "package.zip")
	err = pariseadata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pariseadata_bin_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata.tgz"
	pariseadata_cmd_bin := exec.Command("curl", "-L", pariseadata_bin_url, "-o", "binary.bin")
	err = pariseadata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pariseadata_src_url := "https://pari.math.u-bordeaux.fr/pub/pari/packages/seadata.tgz"
	pariseadata_cmd_src := exec.Command("curl", "-L", pariseadata_src_url, "-o", "source.tar.gz")
	err = pariseadata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pariseadata_cmd_direct := exec.Command("./binary")
	err = pariseadata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pari")
exec.Command("latte", "install", "pari")
}
