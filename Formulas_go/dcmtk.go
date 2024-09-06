package main

// Dcmtk - OFFIS DICOM toolkit command-line utilities
// Homepage: https://dicom.offis.de/dcmtk.php.en

import (
	"fmt"
	
	"os/exec"
)

func installDcmtk() {
	// Método 1: Descargar y extraer .tar.gz
	dcmtk_tar_url := "https://dicom.offis.de/download/dcmtk/dcmtk368/dcmtk-3.6.8.tar.gz"
	dcmtk_cmd_tar := exec.Command("curl", "-L", dcmtk_tar_url, "-o", "package.tar.gz")
	err := dcmtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcmtk_zip_url := "https://dicom.offis.de/download/dcmtk/dcmtk368/dcmtk-3.6.8.zip"
	dcmtk_cmd_zip := exec.Command("curl", "-L", dcmtk_zip_url, "-o", "package.zip")
	err = dcmtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcmtk_bin_url := "https://dicom.offis.de/download/dcmtk/dcmtk368/dcmtk-3.6.8.bin"
	dcmtk_cmd_bin := exec.Command("curl", "-L", dcmtk_bin_url, "-o", "binary.bin")
	err = dcmtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcmtk_src_url := "https://dicom.offis.de/download/dcmtk/dcmtk368/dcmtk-3.6.8.src.tar.gz"
	dcmtk_cmd_src := exec.Command("curl", "-L", dcmtk_src_url, "-o", "source.tar.gz")
	err = dcmtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcmtk_cmd_direct := exec.Command("./binary")
	err = dcmtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
