package main

// Cdi - C and Fortran Interface to access Climate and NWP model Data
// Homepage: https://code.mpimet.mpg.de/projects/cdi

import (
	"fmt"
	
	"os/exec"
)

func installCdi() {
	// Método 1: Descargar y extraer .tar.gz
	cdi_tar_url := "https://code.mpimet.mpg.de/attachments/download/29309/cdi-2.4.0.tar.gz"
	cdi_cmd_tar := exec.Command("curl", "-L", cdi_tar_url, "-o", "package.tar.gz")
	err := cdi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdi_zip_url := "https://code.mpimet.mpg.de/attachments/download/29309/cdi-2.4.0.zip"
	cdi_cmd_zip := exec.Command("curl", "-L", cdi_zip_url, "-o", "package.zip")
	err = cdi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdi_bin_url := "https://code.mpimet.mpg.de/attachments/download/29309/cdi-2.4.0.bin"
	cdi_cmd_bin := exec.Command("curl", "-L", cdi_bin_url, "-o", "binary.bin")
	err = cdi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdi_src_url := "https://code.mpimet.mpg.de/attachments/download/29309/cdi-2.4.0.src.tar.gz"
	cdi_cmd_src := exec.Command("curl", "-L", cdi_src_url, "-o", "source.tar.gz")
	err = cdi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdi_cmd_direct := exec.Command("./binary")
	err = cdi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: eccodes")
	exec.Command("latte", "install", "eccodes").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
	fmt.Println("Instalando dependencia: netcdf")
	exec.Command("latte", "install", "netcdf").Run()
	fmt.Println("Instalando dependencia: proj")
	exec.Command("latte", "install", "proj").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
