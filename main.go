package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func printEULA() bool {
	const eula = `"THIS END-USER LICENSE AGREEMENT ("EULA") is a legal agreement between you ("you","your") (either an individual or a single entity) and Music Tribe IP Limited ("MG-IP"), for the software that accompanies this EULA, which includes associated media and MG-IP Internet-based services ("Software"). Any amendment or addendum to this EULA may accompany the Software. YOU AGREE TO BE BOUND BY THE TERMS OF THIS EULA BY INSTALLING, COPYING, OR USING THE SOFTWARE. IF YOU DO NOT AGREE, DO NOT INSTALL, COPY, OR USE THE SOFTWARE. YOU MAY RETURN IT TO YOUR PLACE OF PURCHASE FOR A FULL REFUND, IF APPLICABLE."

	1. GENERAL
	
	1.1. The Software, documentation, interfaces, content, fonts and any data that accompanied your Product ("Original Software"), as may be updated or replaced by feature enhancements, software updates, supplements, add-on components, internet-based services components or system restore software provided by MG-IP ("Software Updates"), whether in read-only memory, on any other media or in any other form (the Original Software and Software Updates are collectively referred to as the "Software") are licensed, not sold, to you by MG-IP for use only under the terms of this EULA. MG-IP and its licensors retain ownership of the Software itself and reserve all rights not expressly granted to you. MG-IP or its suppliers own the title, copyright and other intellectual property rights in the Software. This EULA does not grant you any rights to trademarks or service marks of MG-IP.
	
	1.2. MG-IP, at its discretion, may provide future Software Updates for your Product. The Software Updates, if any, may not necessarily include all existing software features or new features that MG-IP releases for newer or other models of the Product. The terms of this EULA will govern any Software Updates provided by MG-IP that replace and/or supplement the Original Software, unless a separate license accompanies the Software Update, in which case the terms of that license will govern.
	
	2. GRANT OF LICENSE
	
	2.1. Subject to the terms and conditions of this EULA, MG-IP hereby grants you a limited, non-exclusive license to use the Software on a single Product that you own or control. Except as permitted in Section 2.2 below, and unless as provided in a separate agreement between you and MG-IP, this EULA does not allow the Software to exist on more than one Product at a time, and you may not distribute or make the Software available over a network where it could be used by multiple devices at the same time. You may not rent, lease, lend, sell, redistribute or sublicense the Software.
	
	2.2. Subject to the terms and conditions of this EULA, MG-IP grants you a limited non-exclusive license to download any Software Updates that MG-IP makes available for your model of the Product to update or restore the Software on any such Product that you own or control. This EULA does not allow you to update or restore any Product that you do not control or own, and you may not distribute or make the Software Updates available over a network where they could be used by multiple devices at the same time. If you download a Software Update to your computer, you may make one copy of the Software Updates stored on your computer in machine-readable form for backup purposes only, provided that the backup copy must include all copyright or other proprietary notices contained on the original.
	
	2.3. You acknowledge and agree that you may not enable others to copy (except as expressly permitted by this EULA), decompile, reverse engineer, disassemble, attempt to derive the source code of, decrypt, modify or create derivative works of the Software or any services provided by the Software, or any part thereof (except as and only to the extent any foregoing restriction is prohibited by applicable law or to the extent as may be permitted by licensing terms governing use of open-sourced components included with the Software). Any attempt to do so is a violation of the rights of MG-IP.
	
	2.4. By storing content on your Product, you make a digital copy. In some jurisdictions, it is unlawful to make digital copies without prior permission from the rights holder. You may use the Software to reproduce materials so long as such use is limited to reproduction of non-copyrighted materials, materials in which you own the copyright, or materials in which you have authorization or legal permission to reproduce.
	
	2.5. You agree to use the Software in compliance with all applicable laws, including local laws of the country or region in which you reside or in which you download or use the Software.
	
	2.6. Use of and access to certain features of the Software may require you to apply for a unique user name and password combination.
	
	2.7. You may not be able to exercise your rights to the Software under this EULA after a finite number of product launches unless you activate your copy of the Software in the manner described during the launch sequence. You may also need to reactivate the Software if you modify your Product or alter the Software. MG-IP will use those measures to confirm you have a legally licensed copy of the Software. If you do not use a licensed copy of the Software, you may not install the Software or future Software Updates. MG-IP will not collect any personally identifiable information from your device during this process.
	
	2.8. The documentation that accompanies the Software is licensed for internal, noncommercial reference purposes only.
	
	3. TRANSFER:
	
	You may not rent, lease, lend, sell, redistribute, sublicense or provide commercial hosting services with the Software. You may, however, make a one-time permanent transfer of all of your license rights to the Software to another end user in connection with the transfer of ownership of your Product, provided that: (i) the transfer must include your Product and all of the Software, including all its component parts, original media, printed materials and this EULA; (ii) you do not retain any copies of the Software, full or partial, including copies stored on a computer or other storage device; and (iii) the end user receiving the Software reads and agrees to accept the terms and conditions of this EULA.
	
	4. CONSENT TO USE OF DATA:
	
	You agree that MG-IP and its affiliates may collect, maintain, process and use diagnostic, technical, usage and related information gathered as part of the product support services provided to you, if any, related to the Software, and to verify compliance with the terms of this EULA. MG-IP may use this information solely to improve its products or to provide customized services or technologies to you and will not disclose this information in a form that personally identifies you.
	
	5. CONSENT TO USE OF DATA:
	
	5.1. To use Software identified as an upgrade, you must first be licensed for the Software identified by MG-IP as eligible for the upgrade. After installation of the upgrade, you may no longer use the Original Software that formed the basis for your upgrade eligibility, except as part of the upgraded Software.
	
	5.2. MG-IP, at its discretion, may provide future Software Updates for your Product. The Software Updates, if any, may not include all existing software features or new features that MG-IP releases for newer or other models of Products. The terms of this EULA will govern any Software Updates provided by MG-IP that replace and/or supplement the Original Software, unless such Software Update is accompanied by a separate license in which case the terms of that license will govern.
	
	6. SEPARATION OF COMPONENTS:
	
	MG-IP licenses the Software as a single product. Its component parts may not be separated for use on more than one Product.
	
	7. NOT FOR RESALE SOFTWARE:
	
	Software identified as "Not for Resale" or "NFR" may not be sold or otherwise transferred for value, or used for any purpose other than demonstration, test or evaluation.
	
	8. SUPPORT SERVICES:
	
	MG-IP provides the Software "as is" and may not provide support services for it.
	
	9. DIGITAL CERTIFICATES:
	
	The Software contains functionality that allows it to accept digital certificates issued from MG-IP or from third parties. YOU ARE SOLELY RESPONSIBLE FOR THE DECISION OF WHETHER OR NOT TO RELY ON A CERTIFICATE WHETHER ISSUED BY MG-IP OR A THIRD PARTY. YOUR USE OF DIGITAL CERTIFICATES IS AT YOUR SOLE RISK. TO THE MAXIMUM EXTENT PERMITTED BY APPLICABLE LAW, MG-IP MAKES NO WARRANTIES OR REPRESENTATIONS, EXPRESS OR IMPLIED, AS TO MERCHANTABILITY OR FITNESS FOR ANY PARTICULAR PURPOSE, ACCURACY, SECURITY OR NONINFRINGEMENT OF THIRD PARTY RIGHTS WITH RESPECT TO DIGITAL CERTIFICATES.
	
	10. EXPORT RESTRICTIONS:
	
	The Software is subject to export laws and regulations. You agree to comply with all applicable international and national laws that apply to the Software, including the export regulations, as well as end-user, end-use, and destination restrictions issued by the Grand Duchy of Luxembourg and other governments.
	
	11. TERMINATION:
	
	This EULA is effective until terminated. Your rights under this EULA will terminate automatically or otherwise cease to be effective without notice from MG-IP if you fail to comply with any term(s) of this EULA. Upon the termination of this EULA, you shall cease all use of the Software and destroy all copies of the Software and all of its component parts. Sections 9, 11, 12, 13, 14, and 16 of this EULA shall survive any such termination.
	
	12. DISCLAIMER OF WARRANTIES
	
	12.1. YOU EXPRESSLY ACKNOWLEDGE AND AGREE THAT, TO THE EXTENT PERMITTED BY APPLICABLE LAW, USE OF THE SOFTWARE IS AT YOUR SOLE RISK AND THAT THE ENTIRE RISK AS TO SATISFACTORY QUALITY,PERFORMANCE, ACCURACY AND EFFORT IS WITH YOU.
	
	12.2. TO THE MAXIMUM EXTENT PERMITTED BY APPLICABLE LAW, THE SOFTWARE AND SERVICES PROVIDED BY THE SOFTWARE ARE PROVIDED "AS IS" AND "AS AVAILABLE", WITH ALL FAULTS AND WITHOUT WARRANTY OF ANY KIND, AND MG-IP AND MG-IP'S LICENSORS HEREBY DISCLAIM ALL WARRANTIES AND CONDITIONS WITH RESPECT TO THE SOFTWARE, EITHER EXPRESS, IMPLIED OR STATUTORY, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES AND/OR CONDITIONS OF MERCHANTABILITY, SATISFACTORY QUALITY, FITNESS FOR A PARTICULAR PURPOSE, ACCURACY, QUIET ENJOYMENT, AND NONINFRINGEMENT OF THIRD PARTY RIGHTS.
	
	12.3. MG-IP DOES NOT WARRANT AGAINST INTERFERENCE WITH YOUR ENJOYMENT OF THE SOFTWARE, THAT THE FUNCTIONS CONTAINED IN OR SERVICES PERFORMED BY THE SOFTWARE WILL MEET YOUR REQUIREMENTS, THAT THE OPERATION OF THE SOFTWARE WILL BE UNINTERRUPTED OR ERROR-FREE, THAT ANY SERVICE WILL CONTINUE TO BE MADE AVAILABLE, THAT DEFECTS INTHE SOFTWARE WILL BE CORRECTED OR THAT THE SOFTWARE WILL BE COMPATIBLE OR WORK WITH ANY THIRD PARTY SOFTWARE, APPLICATIONS OR THIRD PARTY SERVICES. INSTALLATION OF THIS SOFTWARE MAY AFFECT THE USABILITY OF THIRD PARTY SOFTWARE, APPLICATIONS OR THIRD PARTY SERVICES.
	
	12.4. YOU FURTHER ACKNOWLEDGE THAT THE SOFTWARE IS NOT INTENDED OR SUITABLE FOR USE IN SITUATIONS OR ENVIRONMENTS WHERE THE FAILURE OR TIME DELAYS OF, OR ERRORS OR INACCURACIES IN, THE CONTENT, DATA OR INFORMATION PROVIDED BY THE SOFTWARE COULD LEAD TO DEATH, PERSONAL INJURY, OR SEVERE PHYSICAL OR ENVIRONMENTAL DAMAGE, INCLUDING WITHOUT LIMITATION THE OPERATION OF NUCLEAR FACILITIES, AIRCRAFT NAVIGATION OR COMMUNICATION SYSTEMS, AIR TRAFFIC CONTROL, LIFE SUPPORT OR WEAPONS SYSTEMS.
	
	12.5. NO ORAL OR WRITTEN INFORMATION OR ADVICE GIVEN BY MG-IP OR AN AUTHORIZED REPRESENTATIVE SHALL CREATE A WARRANTY. SHOULD THE SOFTWARE PROVE DEFECTIVE, YOU ASSUME THE ENTIRE COST OF ALL NECESSARY SERVICING, REPAIR OR CORRECTION. SOME JURISDICTIONS DO NOT ALLOW THE EXCLUSION OF IMPLIED WARRANTIES OR LIMITATIONS ON APPLICABLE STATUTORY RIGHTS OF A CONSUMER AND, THEREFORE, THE ABOVE EXCLUSION AND LIMITATIONS MAY NOT APPLY TO YOU.
	
	13. LIMITATION OF LIABILITY:
	
	TO THE MAXIMUM EXTENT PERMITTED BY APPLICABLE LAW, IN NO EVENT SHALL MG-IP, ITS PARENT, AFFILIATES OR DISTRIBUTORS BE LIABLE FOR PERSONAL INJURY, OR ANY INCIDENTAL, SPECIAL, PUNITIVE, INDIRECT OR CONSEQUENTIAL DAMAGES WHATSOEVER, INCLUDING, WITHOUT LIMITATION, DAMAGES FOR LOSS OF PROFITS OR CONFIDENTIAL OR OTHER INFORMATION, CORRUPTION OR LOSS OF DATA, FAILURE TO TRANSMIT OR RECEIVE ANY DATA, BUSINESS INTERRUPTION OR ANY OTHER COMMERCIAL DAMAGES OR LOSSES, ARISING OUT OF OR RELATED TO YOUR USE OF OR INABILITY TO USE THE SOFTWARE OR ANY THIRD PARTY SOFTWARE OR APPLICATIONS IN CONJUNCTION WITH THE SOFTWARE, HOWEVER CAUSED, REGARDLESS OF THE THEORY OF LIABILITY (CONTRACT, TORT OR OTHERWISE) AND EVEN IF MG-IP HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGES. SOME JURISDICTIONS DO NOT ALLOW THE LIMITATION OF LIABILITY FOR PERSONAL INJURY, OR OF INCIDENTAL OR CONSEQUENTIAL DAMAGES AND, THEREFORE, THIS LIMITATION MAY NOT APPLY TO YOU. In no event shall MG-IP's total liability to you for all damages (other than as may be required by applicable law in cases involving personal injury) exceed the amount of Fifty United States Dollars (USD50.00). The foregoing limitations will apply even if the above stated remedy fails of its essential purpose.
	
	14. EXCLUSIVE REMEDIES:
	
	The entire liability of MG-IP, its parent, affiliates and/or distributors and your exclusive remedy for any breach of this EULA or for any other liability relating to the Software shall be, at MG-IP's option, (a) return of the amount paid (if any) for the Software, or (b) repair or replacement of the Software that is returned to MG-IP with a copy of your receipt. You will receive the remedy elected by MG-IP without charge, except that you are responsible for any expenses you may incur (e.g., cost of shipping the Software to MG-IP). However, this remedy is unavailable if failure of the Software resulted from accident, abuse, misapplication, abnormal use or a virus.
	
	15. THIRD PARTY SOFTWARE
	
	Third party software and data ("Third Party Software") may be attached to the Software. You acknowledge and agree that you must abide by the provisions of any Agreement provided with the Third Party Software and that the party providing the Third Party Software is responsible for any warranty or liability related to or arising from the Third Party Software. MG-IP is not responsible in any way for the Third Party Software or your use thereof.
	
	MG-IP provides no express warranties as to the Third Party Software. IN ADDITION, MG-IP EXPRESSLY DISCLAIMS ALL IMPLIED WARRANTIES, INCLUDING BUT NOT LIMITED TO THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE, as to the Third Party Software.
	
	MG-IP is not liable to you or any other person for any damages, including, without limitation, any direct, indirect, incidental or consequential damages, expenses, lost profits, lost data or other damages arising out of the use, misuse or inability to use the Third Party Software.
	
	16. ENTIRE AGREEMENT:
	
	This EULA, and the terms for any supplements, updates, internet-based services and support services that you use, constitute the entire agreement for the Software.
	
	17. APPLICABLE LAW
	
	17.1. The laws of the Grand Duchy of Luxembourg, excluding its conflicts of law rules, govern this EULA and your use of the Software. Your use of the Software may also be subject to other local, state, national or international laws. The applicability of the Uniform Commercial Code (UCC) and any other laws that direct the application of the laws of any other jurisdiction are expressly excluded. Any dispute arising out of or in connection with this EULA shall be referred to and finally resolved by arbitration in Luxembourg by a sole arbitrator with the arbitration process as foreseen by the new code of civil procedure.
	
	17.2. No amendment to or modification of this EULA will bind either party unless in writing and signed by MG-IP. Any translation of this EULA is done for local requirements and in the event of a dispute between the English and any non-English versions, the English version of this EULA shall govern, to the extent not prohibited by local law in your jurisdiction.
	
	18. MISCELLANEOUS:
	
	If any provision(s) contained in this EULA is or becomes invalid, illegal or unenforceable, in whole or in part, such invalidity, illegality, or unenforceability shall not affect the remaining provisions and portions of thereof, and the invalid, illegal, or unenforceable provision(s) shall be deemed modified so as to have the most similar result that is valid and enforceable under applicable Luxembourg law and/or any other relevant applicable law as the case may be.
	
	Copyright © 2019 Music Tribe Global Brands Ltd. All rights reserved. | Privacy Policy | Imprint & Terms of Use`

	log.Println(eula + "\n")

	reader := bufio.NewReader(os.Stdin)
	log.Print("I Agree to the about terms (y/N): ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	a := strings.ToLower(strings.TrimSuffix(answer, "\n"))

	switch {
	case a == "y":
		return true
	case a == "yes":
		return true
	case a == "n":
		return false
	case a == "no":
		return false
	default:
		log.Fatalf("unacceptable input '%s': only \"Y\" or \"N\" is accepted here", a)
		return false
	}
}

func download(version string) error {
	var (
		URL         string
		outFileName string
	)

	switch runtime.GOOS {
	case "linux":
		switch strings.ToLower(runtime.GOARCH) {
		case "arm":
			URL = fmt.Sprintf("https://downloads.music-group.com/software/behringer/XAIR/X-AIR-Edit_RASPI_V%s.tar.gz", version)
			outFileName = fmt.Sprintf("X-AIR-Edit_RASPI_V%s.tar.gz", version)
		case "arm64":
			URL = fmt.Sprintf("https://downloads.music-group.com/software/behringer/XAIR/X-AIR-Edit_RASPI_V%s.tar.gz", version)
			outFileName = fmt.Sprintf("X-AIR-Edit_RASPI_V%s.tar.gz", version)
		case "amd64":
			URL = fmt.Sprintf("https://downloads.music-group.com/software/behringer/XAIR/X-AIR-Edit_LINUX_X64_V%s.tar.gz", version)
			outFileName = fmt.Sprintf("X-AIR-Edit_LINUX_X64_V%s.tar.gz", version)
		case "386":
			URL = fmt.Sprintf("https://downloads.music-group.com/software/behringer/XAIR/X-AIR-Edit_LINUX_V%s.tar.gz", version)
			outFileName = fmt.Sprintf("X-AIR-Edit_LINUX_V%s.tar.gz", version)
		default:
			return fmt.Errorf("unsupported linux processor architecture: '%s'", runtime.GOARCH)
		}
	case "darwin":
		URL = fmt.Sprintf("https://downloads.music-group.com/software/behringer/XAIR/X-AIR-Edit_MAC_V%s.zip", version)
		outFileName = fmt.Sprintf("X-AIR-Edit_MAC_V%s.zip", version)
	case "windows":
		URL = fmt.Sprintf("https://downloads.music-group.com/software/behringer/XAIR/X-AIR-Edit_PC_V%s.zip", version)
		outFileName = fmt.Sprintf("X-AIR-Edit_PC_V%s.zip", version)
	default:
		return fmt.Errorf("unsupported operating system: '%s'", runtime.GOOS)
	}

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error fetching from %s", URL)
	}

	f, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func execute() {
	var (
		acceptEULA bool
		version    string

		rootCmd = &cobra.Command{
			Use:     "xair-downloader",
			Version: "0.1.0",
			Short:   "A Downloader for the Behringer X-AIR program",
			Long:    "",
			Example: "",
			Run: func(cmd *cobra.Command, args []string) {
				if !acceptEULA {
					if !printEULA() {
						log.Fatal("license not accepted, cancelling")
					}
				}

				log.Println("starting download")
				if err := download(version); err != nil {
					log.Fatal(err)
				}
				defer log.Println("done")
			},
		}
	)

	rootCmd.Flags().BoolVarP(&acceptEULA, "accept-eula", "y", false, "automatically accept the Music Tribe End-User License Agreement")
	rootCmd.Flags().StringVarP(&version, "xair-version", "v", "1.5", "the version of X-AIR to download")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execute()
}
