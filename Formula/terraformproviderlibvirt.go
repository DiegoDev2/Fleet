package main

// TerraformProviderLibvirt - Terraform provisioning with Linux KVM using libvirt
// Homepage: https://github.com/dmacvicar/terraform-provider-libvirt

import (
	"fmt"
	
	"os/exec"
)

func installTerraformProviderLibvirt() {
	// Método 1: Descargar y extraer .tar.gz
	terraformproviderlibvirt_tar_url := "https://github.com/dmacvicar/terraform-provider-libvirt/archive/refs/tags/v0.7.6.tar.gz"
	terraformproviderlibvirt_cmd_tar := exec.Command("curl", "-L", terraformproviderlibvirt_tar_url, "-o", "package.tar.gz")
	err := terraformproviderlibvirt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformproviderlibvirt_zip_url := "https://github.com/dmacvicar/terraform-provider-libvirt/archive/refs/tags/v0.7.6.zip"
	terraformproviderlibvirt_cmd_zip := exec.Command("curl", "-L", terraformproviderlibvirt_zip_url, "-o", "package.zip")
	err = terraformproviderlibvirt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformproviderlibvirt_bin_url := "https://github.com/dmacvicar/terraform-provider-libvirt/archive/refs/tags/v0.7.6.bin"
	terraformproviderlibvirt_cmd_bin := exec.Command("curl", "-L", terraformproviderlibvirt_bin_url, "-o", "binary.bin")
	err = terraformproviderlibvirt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformproviderlibvirt_src_url := "https://github.com/dmacvicar/terraform-provider-libvirt/archive/refs/tags/v0.7.6.src.tar.gz"
	terraformproviderlibvirt_cmd_src := exec.Command("curl", "-L", terraformproviderlibvirt_src_url, "-o", "source.tar.gz")
	err = terraformproviderlibvirt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformproviderlibvirt_cmd_direct := exec.Command("./binary")
	err = terraformproviderlibvirt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libvirt")
	exec.Command("latte", "install", "libvirt").Run()
}
