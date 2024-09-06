package main

// Mkhexgrid - Fully-configurable hex grid generator
// Homepage: https://www.nomic.net/~uckelman/mkhexgrid/

import (
	"fmt"
	
	"os/exec"
)

func installMkhexgrid() {
	// Método 1: Descargar y extraer .tar.gz
	mkhexgrid_tar_url := "https://www.nomic.net/~uckelman/mkhexgrid/releases/mkhexgrid-0.1.1.src.tar.bz2"
	mkhexgrid_cmd_tar := exec.Command("curl", "-L", mkhexgrid_tar_url, "-o", "package.tar.gz")
	err := mkhexgrid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkhexgrid_zip_url := "https://www.nomic.net/~uckelman/mkhexgrid/releases/mkhexgrid-0.1.1.src.tar.bz2"
	mkhexgrid_cmd_zip := exec.Command("curl", "-L", mkhexgrid_zip_url, "-o", "package.zip")
	err = mkhexgrid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkhexgrid_bin_url := "https://www.nomic.net/~uckelman/mkhexgrid/releases/mkhexgrid-0.1.1.src.tar.bz2"
	mkhexgrid_cmd_bin := exec.Command("curl", "-L", mkhexgrid_bin_url, "-o", "binary.bin")
	err = mkhexgrid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkhexgrid_src_url := "https://www.nomic.net/~uckelman/mkhexgrid/releases/mkhexgrid-0.1.1.src.tar.bz2"
	mkhexgrid_cmd_src := exec.Command("curl", "-L", mkhexgrid_src_url, "-o", "source.tar.gz")
	err = mkhexgrid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkhexgrid_cmd_direct := exec.Command("./binary")
	err = mkhexgrid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
}
