package main

// Eccodes - Decode and encode messages in the GRIB 1/2 and BUFR 3/4 formats
// Homepage: https://confluence.ecmwf.int/display/ECC

import (
	"fmt"
	
	"os/exec"
)

func installEccodes() {
	// Método 1: Descargar y extraer .tar.gz
	eccodes_tar_url := "https://confluence.ecmwf.int/download/attachments/45757960/eccodes-2.37.0-Source.tar.gz"
	eccodes_cmd_tar := exec.Command("curl", "-L", eccodes_tar_url, "-o", "package.tar.gz")
	err := eccodes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eccodes_zip_url := "https://confluence.ecmwf.int/download/attachments/45757960/eccodes-2.37.0-Source.zip"
	eccodes_cmd_zip := exec.Command("curl", "-L", eccodes_zip_url, "-o", "package.zip")
	err = eccodes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eccodes_bin_url := "https://confluence.ecmwf.int/download/attachments/45757960/eccodes-2.37.0-Source.bin"
	eccodes_cmd_bin := exec.Command("curl", "-L", eccodes_bin_url, "-o", "binary.bin")
	err = eccodes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eccodes_src_url := "https://confluence.ecmwf.int/download/attachments/45757960/eccodes-2.37.0-Source.src.tar.gz"
	eccodes_cmd_src := exec.Command("curl", "-L", eccodes_src_url, "-o", "source.tar.gz")
	err = eccodes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eccodes_cmd_direct := exec.Command("./binary")
	err = eccodes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: netcdf")
	exec.Command("latte", "install", "netcdf").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
}
