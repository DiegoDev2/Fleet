package main

// Ncview - Visual browser for netCDF format files
// Homepage: https://cirrus.ucsd.edu/ncview/

import (
	"fmt"
	
	"os/exec"
)

func installNcview() {
	// Método 1: Descargar y extraer .tar.gz
	ncview_tar_url := "https://cirrus.ucsd.edu/~pierce/ncview/ncview-2.1.10.tar.gz"
	ncview_cmd_tar := exec.Command("curl", "-L", ncview_tar_url, "-o", "package.tar.gz")
	err := ncview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncview_zip_url := "https://cirrus.ucsd.edu/~pierce/ncview/ncview-2.1.10.zip"
	ncview_cmd_zip := exec.Command("curl", "-L", ncview_zip_url, "-o", "package.zip")
	err = ncview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncview_bin_url := "https://cirrus.ucsd.edu/~pierce/ncview/ncview-2.1.10.bin"
	ncview_cmd_bin := exec.Command("curl", "-L", ncview_bin_url, "-o", "binary.bin")
	err = ncview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncview_src_url := "https://cirrus.ucsd.edu/~pierce/ncview/ncview-2.1.10.src.tar.gz"
	ncview_cmd_src := exec.Command("curl", "-L", ncview_src_url, "-o", "source.tar.gz")
	err = ncview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncview_cmd_direct := exec.Command("./binary")
	err = ncview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libice")
exec.Command("latte", "install", "libice")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsm")
exec.Command("latte", "install", "libsm")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxaw")
exec.Command("latte", "install", "libxaw")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: udunits")
exec.Command("latte", "install", "udunits")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
}
