package main

// Dcraw - Digital camera RAW photo decoding software
// Homepage: https://www.dechifro.org/dcraw/

import (
	"fmt"
	
	"os/exec"
)

func installDcraw() {
	// Método 1: Descargar y extraer .tar.gz
	dcraw_tar_url := "https://www.dechifro.org/dcraw/archive/dcraw-9.28.0.tar.gz"
	dcraw_cmd_tar := exec.Command("curl", "-L", dcraw_tar_url, "-o", "package.tar.gz")
	err := dcraw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcraw_zip_url := "https://www.dechifro.org/dcraw/archive/dcraw-9.28.0.zip"
	dcraw_cmd_zip := exec.Command("curl", "-L", dcraw_zip_url, "-o", "package.zip")
	err = dcraw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcraw_bin_url := "https://www.dechifro.org/dcraw/archive/dcraw-9.28.0.bin"
	dcraw_cmd_bin := exec.Command("curl", "-L", dcraw_bin_url, "-o", "binary.bin")
	err = dcraw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcraw_src_url := "https://www.dechifro.org/dcraw/archive/dcraw-9.28.0.src.tar.gz"
	dcraw_cmd_src := exec.Command("curl", "-L", dcraw_src_url, "-o", "source.tar.gz")
	err = dcraw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcraw_cmd_direct := exec.Command("./binary")
	err = dcraw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
}
