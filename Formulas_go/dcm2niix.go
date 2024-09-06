package main

// Dcm2niix - DICOM to NIfTI converter
// Homepage: https://www.nitrc.org/plugins/mwiki/index.php/dcm2nii:MainPage

import (
	"fmt"
	
	"os/exec"
)

func installDcm2niix() {
	// Método 1: Descargar y extraer .tar.gz
	dcm2niix_tar_url := "https://github.com/rordenlab/dcm2niix/archive/refs/tags/v1.0.20240202.tar.gz"
	dcm2niix_cmd_tar := exec.Command("curl", "-L", dcm2niix_tar_url, "-o", "package.tar.gz")
	err := dcm2niix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcm2niix_zip_url := "https://github.com/rordenlab/dcm2niix/archive/refs/tags/v1.0.20240202.zip"
	dcm2niix_cmd_zip := exec.Command("curl", "-L", dcm2niix_zip_url, "-o", "package.zip")
	err = dcm2niix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcm2niix_bin_url := "https://github.com/rordenlab/dcm2niix/archive/refs/tags/v1.0.20240202.bin"
	dcm2niix_cmd_bin := exec.Command("curl", "-L", dcm2niix_bin_url, "-o", "binary.bin")
	err = dcm2niix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcm2niix_src_url := "https://github.com/rordenlab/dcm2niix/archive/refs/tags/v1.0.20240202.src.tar.gz"
	dcm2niix_cmd_src := exec.Command("curl", "-L", dcm2niix_src_url, "-o", "source.tar.gz")
	err = dcm2niix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcm2niix_cmd_direct := exec.Command("./binary")
	err = dcm2niix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
