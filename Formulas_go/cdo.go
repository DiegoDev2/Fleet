package main

// Cdo - Climate Data Operators
// Homepage: https://code.mpimet.mpg.de/projects/cdo

import (
	"fmt"
	
	"os/exec"
)

func installCdo() {
	// Método 1: Descargar y extraer .tar.gz
	cdo_tar_url := "https://code.mpimet.mpg.de/attachments/download/29616/cdo-2.4.3.tar.gz"
	cdo_cmd_tar := exec.Command("curl", "-L", cdo_tar_url, "-o", "package.tar.gz")
	err := cdo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdo_zip_url := "https://code.mpimet.mpg.de/attachments/download/29616/cdo-2.4.3.zip"
	cdo_cmd_zip := exec.Command("curl", "-L", cdo_zip_url, "-o", "package.zip")
	err = cdo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdo_bin_url := "https://code.mpimet.mpg.de/attachments/download/29616/cdo-2.4.3.bin"
	cdo_cmd_bin := exec.Command("curl", "-L", cdo_bin_url, "-o", "binary.bin")
	err = cdo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdo_src_url := "https://code.mpimet.mpg.de/attachments/download/29616/cdo-2.4.3.src.tar.gz"
	cdo_cmd_src := exec.Command("curl", "-L", cdo_src_url, "-o", "source.tar.gz")
	err = cdo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdo_cmd_direct := exec.Command("./binary")
	err = cdo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: eccodes")
exec.Command("latte", "install", "eccodes")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
